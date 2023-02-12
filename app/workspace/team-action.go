package workspace

import (
	"github.com/pkg/errors"

	"github.com/kyleu/rituals/app/action"
	"github.com/kyleu/rituals/app/enum"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
)

func (s *Service) ActionTeam(p *Params) (*FullTeam, string, string, error) {
	lp := NewLoadParams(p.Ctx, p.Slug, p.Profile, p.Accounts, nil, nil, p.Logger)
	ft, err := p.Svc.LoadTeam(lp)
	if err != nil {
		return nil, "", "", err
	}

	switch p.Act {
	case action.ActUpdate:
		return teamUpdate(p, ft)
	case action.ActMemberUpdate:
		return teamMemberUpdate(p, ft)
	case action.ActMemberRemove:
		return teamMemberRemove(p, ft)
	case action.ActMemberSelf:
		return teamUpdateSelf(p, ft)
	case action.ActComment:
		return teamComment(p, ft)
	case "":
		return nil, "", "", errors.New("field [action] is required")
	default:
		return nil, "", "", errors.Errorf("invalid action [%s]", p.Act)
	}
}

func teamUpdate(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	tgt := ft.Team.Clone()
	tgt.Title = p.Frm.GetStringOpt("title")
	if tgt.Title == "" {
		return nil, "", "", errors.New("title may not be empty")
	}
	tgt.Slug = p.Frm.GetStringOpt("slug")
	if tgt.Slug == "" {
		tgt.Slug = util.Slugify(tgt.Title)
	}
	tx, err := p.Svc.db.StartTransaction(p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	defer func() {
		_ = tx.Rollback()
	}()
	tgt.Slug = p.Svc.t.Slugify(p.Ctx, tgt.ID, tgt.Slug, p.Slug, p.Svc.th, tx, p.Logger)
	tgt.Icon = p.Frm.GetStringOpt("icon")
	tgt.Icon = tgt.IconSafe()
	perms := loadPermissionsForm(p.Frm)
	modelChanged := len(ft.Team.Diff(tgt)) > 0
	permsChanged := len(ft.Permissions.ToPermissions().Diff(perms)) > 0
	if !modelChanged && !permsChanged {
		return ft, MsgNoChangesNeeded, ft.Team.PublicWebPath(), nil
	}
	if modelChanged {
		model, err := p.Svc.SaveTeam(p.Ctx, tgt, ft.Self.UserID, tx, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
		ft.Team = model
		err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActUpdate, model, &ft.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	if permsChanged {
		if err := p.Svc.tp.DeleteWhere(p.Ctx, tx, "team_id = $1", len(ft.Permissions), p.Logger, tgt.ID); err != nil {
			return nil, "", "", err
		}
		newPerms := make(tpermission.TeamPermissions, 0, len(perms))
		for _, x := range perms {
			newPerms = append(newPerms, &tpermission.TeamPermission{TeamID: tgt.ID, Key: x.Key, Value: x.Value})
		}
		if err = p.Svc.tp.Save(p.Ctx, tx, p.Logger, newPerms...); err != nil {
			return nil, "", "", err
		}
		err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActPermissions, perms, &ft.Self.UserID, p.Logger)
		if err != nil {
			return nil, "", "", err
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, "", "", err
	}
	return ft, "Team updated", ft.Team.PublicWebPath(), nil
}

func teamMemberUpdate(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	if ft.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this team")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	role := p.Frm.GetStringOpt("role")
	if role == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	curr := ft.Members.Get(ft.Team.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this team", userID.String())
	}
	curr.Role = enum.MemberStatus(role)
	err := p.Svc.tm.Update(p.Ctx, nil, curr, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberUpdate, curr, &ft.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return ft, MsgMemberUpdated, ft.Team.PublicWebPath(), nil
}

func teamMemberRemove(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	if ft.Self.Role != enum.MemberStatusOwner {
		return nil, "", "", errors.New("you are not the owner of this team")
	}
	userID, _ := p.Frm.GetUUID("userID", false)
	if userID == nil {
		return nil, "", "", errors.New("must provide [userID]")
	}
	curr := ft.Members.Get(ft.Team.ID, *userID)
	if curr == nil {
		return nil, "", "", errors.Errorf("user [%s] is not a member of this team", userID.String())
	}
	err := p.Svc.tm.Delete(p.Ctx, nil, curr.TeamID, curr.UserID, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberRemove, userID, &ft.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return ft, MsgMemberRemoved, ft.Team.PublicWebPath(), nil
}

func teamUpdateSelf(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	choice := p.Frm.GetStringOpt("choice")
	name := p.Frm.GetStringOpt("name")
	if name == "" {
		return nil, "", "", errors.New("must provide [name]")
	}
	ft.Self.Name = name
	err := p.Svc.tm.Update(p.Ctx, nil, ft.Self, p.Logger)
	if err != nil {
		return nil, "", "", err
	}
	if choice == KeyGlobal {
		return nil, "", "", errors.New("can't change global name yet")
	}
	arg := util.ValueMap{"userID": ft.Self.UserID, "name": name}
	err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActMemberUpdate, arg, &ft.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return ft, MsgProfileEdited, ft.Team.PublicWebPath(), nil
}

func teamComment(p *Params, ft *FullTeam) (*FullTeam, string, string, error) {
	if ft.Self == nil {
		return nil, "", "", errors.New("you are not a member of this team")
	}
	c, u, err := commentFromForm(p.Frm, ft.Self.UserID)
	if err != nil {
		return nil, "", "", err
	}
	switch c.Svc {
	case enum.ModelServiceTeam:
		if c.ModelID != ft.Team.ID {
			return nil, "", "", errors.New("this comment refers to a different team")
		}
	case enum.ModelServiceSprint:
		if curr := ft.Sprints.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an sprint that isn't part of this team")
		}
	case enum.ModelServiceEstimate:
		if curr := ft.Estimates.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an estimate that isn't part of this team")
		}
	case enum.ModelServiceStandup:
		if curr := ft.Standups.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an standup that isn't part of this team")
		}
	case enum.ModelServiceRetro:
		if curr := ft.Retros.Get(c.ModelID); curr == nil {
			return nil, "", "", errors.New("this comment refers to an retro that isn't part of this team")
		}
	default:
		return nil, "", "", errors.Errorf("can't comment on object of type [%s]", c.Svc)
	}
	err = p.Svc.c.Save(p.Ctx, nil, p.Logger, c)
	if err != nil {
		return nil, "", "", err
	}
	err = p.Svc.send(enum.ModelServiceTeam, ft.Team.ID, action.ActComment, c, &ft.Self.UserID, p.Logger, p.ConnIDs...)
	if err != nil {
		return nil, "", "", err
	}
	return ft, MsgCommentAdded, ft.Team.PublicWebPath() + u, nil
}
