package web

import (
	"fmt"
	"github.com/gofrs/uuid"
	"net/http"
	"net/url"

	"github.com/kyleu/rituals.dev/app/config"
	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"logur.dev/logur"
)

type RequestContext struct {
	App         *config.AppInfo
	Logger      logur.Logger
	Profile     *util.UserProfile
	Routes      *mux.Router
	Request     *url.URL
	Title       string
	Breadcrumbs Breadcrumbs
	Flashes     []string
	Session     sessions.Session
}

func (r *RequestContext) Route(act string, pairs ...string) string {
	route := r.Routes.Get(act)
	if route == nil {
		r.App.Logger.Warn("cannot find route at path [" + act + "]")
		return "/routenotfound"
	}
	u, err := route.URL(pairs...)
	if err != nil {
		r.App.Logger.Warn("cannot bind route at path [" + act + "]")
		return "/routeerror"
	}
	return u.Path
}

func ExtractContext(w http.ResponseWriter, r *http.Request) RequestContext {
	ai, ok := r.Context().Value("info").(*config.AppInfo)
	if !ok {
		ai.Logger.Warn("cannot load AppInfo")
	}
	routes, ok := r.Context().Value("routes").(*mux.Router)
	if !ok {
		ai.Logger.Warn("cannot load Router")
	}
	session, err := store.Get(r, sessionName)
	if err != nil {
		session = sessions.NewSession(store, sessionName)
	}

	var userID uuid.UUID
	userIDValue, ok := session.Values[util.KeyUser]
	if ok {
		userID, err = uuid.FromString(userIDValue.(string))
		if err != nil {
			ai.Logger.Warn(fmt.Sprintf("cannot parse uuid [%v]: %+v", userIDValue, err))
		}
	} else {
		userID = util.UUID()
		_, err := ai.User.New(userID)
		if err != nil {
			ai.Logger.Warn(fmt.Sprintf("cannot save user: %+v", err))
		}
		session.Values[util.KeyUser] = userID.String()
		err = session.Save(r, w)
		if err != nil {
			ai.Logger.Warn(fmt.Sprintf("cannot save session: %+v", err))
		}
	}

	user, err := ai.User.GetByID(userID, true)
	if err != nil {
		ai.Logger.Warn(fmt.Sprintf("unable to load user profile: %+v", err))
	}
	var prof *util.UserProfile
	if user == nil {
		prof = util.NewUserProfile(userID)
	} else {
		prof = user.ToProfile()
	}

	flashes := make([]string, 0)
	for _, f := range session.Flashes() {
		flashes = append(flashes, fmt.Sprint(f))
	}

	logger := logur.WithFields(ai.Logger, map[string]interface{}{"path": r.URL.Path, "method": r.Method})

	return RequestContext{
		App:         ai,
		Logger:      logger,
		Profile:     prof,
		Routes:      routes,
		Request:     r.URL,
		Title:       util.AppName,
		Breadcrumbs: nil,
		Flashes:     flashes,
		Session:     *session,
	}
}
