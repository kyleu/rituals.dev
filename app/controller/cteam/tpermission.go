// Package cteam - Content managed by Project Forge, see [projectforge.md] for details.
package cteam

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/team/tpermission"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vteam/vtpermission"
)

func TeamPermissionList(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("tpermission", nil, ps.Logger).Sanitize("tpermission")
		ret, err := as.Services.TeamPermission.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Permissions", ret)
		teamIDsByTeamID := lo.Map(ret, func(x *tpermission.TeamPermission, _ int) uuid.UUID {
			return x.TeamID
		})
		teamsByTeamID, err := as.Services.Team.GetMultiple(ps.Context, nil, nil, ps.Logger, teamIDsByTeamID...)
		if err != nil {
			return "", err
		}
		page := &vtpermission.List{Models: ret, TeamsByTeamID: teamsByTeamID, Params: ps.Params}
		return controller.Render(rc, as, page, ps, "team", "tpermission")
	})
}

func TeamPermissionDetail(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData(ret.TitleString()+" (Permission)", ret)

		teamByTeamID, _ := as.Services.Team.Get(ps.Context, nil, ret.TeamID, ps.Logger)

		return controller.Render(rc, as, &vtpermission.Detail{Model: ret, TeamByTeamID: teamByTeamID}, ps, "team", "tpermission", ret.TitleString()+"**permission")
	})
}

func TeamPermissionCreateForm(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &tpermission.TeamPermission{}
		if string(rc.QueryArgs().Peek("prototype")) == util.KeyRandom {
			ret = tpermission.Random()
		}
		ps.SetTitleAndData("Create [TeamPermission]", ret)
		ps.Data = ret
		return controller.Render(rc, as, &vtpermission.Edit{Model: ret, IsNew: true}, ps, "team", "tpermission", "Create")
	})
}

func TeamPermissionRandom(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := as.Services.TeamPermission.Random(ps.Context, nil, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to find random TeamPermission")
		}
		return ret.WebPath(), nil
	})
}

func TeamPermissionCreate(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamPermission from form")
		}
		err = as.Services.TeamPermission.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created TeamPermission")
		}
		msg := fmt.Sprintf("TeamPermission [%s] created", ret.String())
		return controller.FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func TeamPermissionEditForm(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.SetTitleAndData("Edit "+ret.String(), ret)
		return controller.Render(rc, as, &vtpermission.Edit{Model: ret}, ps, "team", "tpermission", ret.String())
	})
}

func TeamPermissionEdit(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := tpermissionFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse TeamPermission from form")
		}
		frm.TeamID = ret.TeamID
		frm.Key = ret.Key
		frm.Value = ret.Value
		err = as.Services.TeamPermission.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update TeamPermission [%s]", frm.String())
		}
		msg := fmt.Sprintf("TeamPermission [%s] updated", frm.String())
		return controller.FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func TeamPermissionDelete(rc *fasthttp.RequestCtx) {
	controller.Act("tpermission.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := tpermissionFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.TeamPermission.Delete(ps.Context, nil, ret.TeamID, ret.Key, ret.Value, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete permission [%s]", ret.String())
		}
		msg := fmt.Sprintf("TeamPermission [%s] deleted", ret.String())
		return controller.FlashAndRedir(true, msg, "/admin/db/team/permission", rc, ps)
	})
}

func tpermissionFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*tpermission.TeamPermission, error) {
	teamIDArgStr, err := cutil.RCRequiredString(rc, "teamID", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [teamID] as an argument")
	}
	teamIDArgP := util.UUIDFromString(teamIDArgStr)
	if teamIDArgP == nil {
		return nil, errors.Errorf("argument [teamID] (%s) is not a valid UUID", teamIDArgStr)
	}
	teamIDArg := *teamIDArgP
	keyArg, err := cutil.RCRequiredString(rc, "key", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [key] as a string argument")
	}
	valueArg, err := cutil.RCRequiredString(rc, "value", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [value] as a string argument")
	}
	return as.Services.TeamPermission.Get(ps.Context, nil, teamIDArg, keyArg, valueArg, ps.Logger)
}

func tpermissionFromForm(rc *fasthttp.RequestCtx, setPK bool) (*tpermission.TeamPermission, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return tpermission.FromMap(frm, setPK)
}
