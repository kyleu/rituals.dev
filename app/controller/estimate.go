// Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/estimate"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/vestimate"
)

func EstimateList(rc *fasthttp.RequestCtx) {
	Act("estimate.list", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		prms := ps.Params.Get("estimate", nil, ps.Logger).Sanitize("estimate")
		ret, err := as.Services.Estimate.List(ps.Context, nil, prms, ps.Logger)
		if err != nil {
			return "", err
		}
		ps.Title = "Estimates"
		ps.Data = ret
		userIDs := make([]uuid.UUID, 0, len(ret))
		for _, x := range ret {
			userIDs = append(userIDs, x.Owner)
		}
		users, err := as.Services.User.GetMultiple(ps.Context, nil, ps.Logger, userIDs...)
		if err != nil {
			return "", err
		}
		teamIDs := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			teamIDs = append(teamIDs, x.TeamID)
		}
		teams, err := as.Services.Team.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDefererence(teamIDs)...)
		if err != nil {
			return "", err
		}
		sprintIDs := make([]*uuid.UUID, 0, len(ret))
		for _, x := range ret {
			sprintIDs = append(sprintIDs, x.SprintID)
		}
		sprints, err := as.Services.Sprint.GetMultiple(ps.Context, nil, ps.Logger, util.ArrayDefererence(sprintIDs)...)
		if err != nil {
			return "", err
		}
		return Render(rc, as, &vestimate.List{Models: ret, Users: users, Teams: teams, Sprints: sprints, Params: ps.Params}, ps, "estimate")
	})
}

func EstimateDetail(rc *fasthttp.RequestCtx) {
	Act("estimate.detail", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = ret.TitleString() + " (Estimate)"
		ps.Data = ret
		estimateHistoryPrms := ps.Params.Get("estimateHistory", nil, ps.Logger).Sanitize("estimateHistory")
		estimateHistoriesByEstimateID, err := as.Services.EstimateHistory.GetByEstimateID(ps.Context, nil, ret.ID, estimateHistoryPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child histories")
		}
		estimateMemberPrms := ps.Params.Get("estimateMember", nil, ps.Logger).Sanitize("estimateMember")
		estimateMembersByEstimateID, err := as.Services.EstimateMember.GetByEstimateID(ps.Context, nil, ret.ID, estimateMemberPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child members")
		}
		estimatePermissionPrms := ps.Params.Get("estimatePermission", nil, ps.Logger).Sanitize("estimatePermission")
		estimatePermissionsByEstimateID, err := as.Services.EstimatePermission.GetByEstimateID(ps.Context, nil, ret.ID, estimatePermissionPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child permissions")
		}
		storyPrms := ps.Params.Get("story", nil, ps.Logger).Sanitize("story")
		storiesByEstimateID, err := as.Services.Story.GetByEstimateID(ps.Context, nil, ret.ID, storyPrms, ps.Logger)
		if err != nil {
			return "", errors.Wrap(err, "unable to retrieve child stories")
		}
		return Render(rc, as, &vestimate.Detail{
			Model:                           ret,
			Params:                          ps.Params,
			EstimateHistoriesByEstimateID:   estimateHistoriesByEstimateID,
			EstimateMembersByEstimateID:     estimateMembersByEstimateID,
			EstimatePermissionsByEstimateID: estimatePermissionsByEstimateID,
			StoriesByEstimateID:             storiesByEstimateID,
		}, ps, "estimate", ret.String())
	})
}

func EstimateCreateForm(rc *fasthttp.RequestCtx) {
	Act("estimate.create.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := &estimate.Estimate{}
		ps.Title = "Create [Estimate]"
		ps.Data = ret
		return Render(rc, as, &vestimate.Edit{Model: ret, IsNew: true}, ps, "estimate", "Create")
	})
}

func EstimateCreateFormRandom(rc *fasthttp.RequestCtx) {
	Act("estimate.create.form.random", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret := estimate.Random()
		ps.Title = "Create Random Estimate"
		ps.Data = ret
		return Render(rc, as, &vestimate.Edit{Model: ret, IsNew: true}, ps, "estimate", "Create")
	})
}

func EstimateCreate(rc *fasthttp.RequestCtx) {
	Act("estimate.create", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromForm(rc, true)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Estimate from form")
		}
		err = as.Services.Estimate.Create(ps.Context, nil, ps.Logger, ret)
		if err != nil {
			return "", errors.Wrap(err, "unable to save newly-created Estimate")
		}
		msg := fmt.Sprintf("Estimate [%s] created", ret.String())
		return FlashAndRedir(true, msg, ret.WebPath(), rc, ps)
	})
}

func EstimateEditForm(rc *fasthttp.RequestCtx) {
	Act("estimate.edit.form", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		ps.Title = "Edit " + ret.String()
		ps.Data = ret
		return Render(rc, as, &vestimate.Edit{Model: ret}, ps, "estimate", ret.String())
	})
}

func EstimateEdit(rc *fasthttp.RequestCtx) {
	Act("estimate.edit", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		frm, err := estimateFromForm(rc, false)
		if err != nil {
			return "", errors.Wrap(err, "unable to parse Estimate from form")
		}
		frm.ID = ret.ID
		err = as.Services.Estimate.Update(ps.Context, nil, frm, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to update Estimate [%s]", frm.String())
		}
		msg := fmt.Sprintf("Estimate [%s] updated", frm.String())
		return FlashAndRedir(true, msg, frm.WebPath(), rc, ps)
	})
}

func EstimateDelete(rc *fasthttp.RequestCtx) {
	Act("estimate.delete", rc, func(as *app.State, ps *cutil.PageState) (string, error) {
		ret, err := estimateFromPath(rc, as, ps)
		if err != nil {
			return "", err
		}
		err = as.Services.Estimate.Delete(ps.Context, nil, ret.ID, ps.Logger)
		if err != nil {
			return "", errors.Wrapf(err, "unable to delete estimate [%s]", ret.String())
		}
		msg := fmt.Sprintf("Estimate [%s] deleted", ret.String())
		return FlashAndRedir(true, msg, "/estimate", rc, ps)
	})
}

func estimateFromPath(rc *fasthttp.RequestCtx, as *app.State, ps *cutil.PageState) (*estimate.Estimate, error) {
	idArgStr, err := cutil.RCRequiredString(rc, "id", false)
	if err != nil {
		return nil, errors.Wrap(err, "must provide [id] as an argument")
	}
	idArgP := util.UUIDFromString(idArgStr)
	if idArgP == nil {
		return nil, errors.Errorf("argument [id] (%s) is not a valid UUID", idArgStr)
	}
	idArg := *idArgP
	return as.Services.Estimate.Get(ps.Context, nil, idArg, ps.Logger)
}

func estimateFromForm(rc *fasthttp.RequestCtx, setPK bool) (*estimate.Estimate, error) {
	frm, err := cutil.ParseForm(rc)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse form")
	}
	return estimate.FromMap(frm, setPK)
}
