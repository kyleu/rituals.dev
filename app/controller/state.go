// Package controller - Content managed by Project Forge, see [projectforge.md] for details.
package controller

import (
	"fmt"
	"net/http"

	"github.com/kyleu/rituals/app"
	"github.com/kyleu/rituals/app/controller/cutil"
	"github.com/kyleu/rituals/app/util"
	"github.com/kyleu/rituals/views/verror"
)

var (
	_currentAppState       *app.State
	_currentAppRootLogger  util.Logger
	_currentSiteState      *app.State
	_currentSiteRootLogger util.Logger
)

func SetAppState(a *app.State, logger util.Logger) {
	_currentAppState = a
	_currentAppRootLogger = logger
	initApp(a, logger)
}

func SetSiteState(a *app.State, logger util.Logger) {
	_currentSiteState = a
	_currentSiteRootLogger = logger
	initSite(a, logger)
}

func handleError(key string, as *app.State, ps *cutil.PageState, w http.ResponseWriter, r *http.Request, err error) (string, error) {
	w.WriteHeader(http.StatusInternalServerError)

	ps.LogError("error running action [%s]: %+v", key, err)

	if len(ps.Breadcrumbs) == 0 {
		bc := util.StringSplitAndTrim(r.URL.Path, "/")
		bc = append(bc, "Error")
		ps.Breadcrumbs = bc
	}

	if cleanErr := ps.Clean(r, as); cleanErr != nil {
		ps.Logger.Error(cleanErr)
		msg := fmt.Sprintf("error while cleaning request: %+v", cleanErr)
		ps.Logger.Error(msg)
		_, _ = w.Write([]byte(msg))
		return "", cleanErr
	}

	e := util.GetErrorDetail(err, ps.Admin)
	ps.Data = e
	redir, renderErr := Render(w, r, as, &verror.Error{Err: e}, ps)
	if renderErr != nil {
		msg := fmt.Sprintf("error while running error handler: %+v", renderErr)
		ps.Logger.Error(msg)
		_, _ = w.Write([]byte(msg))
		return "", renderErr
	}
	return redir, nil
}
