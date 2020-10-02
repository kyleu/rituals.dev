package transcript

import (
	"fmt"
	"time"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npnweb"
	"github.com/kyleu/rituals.dev/app"

	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
)

type EmailResponse struct {
	Date      *time.Time
	Teams     team.Sessions     `json:"teams"`
	Sprints   sprint.Sessions   `json:"sprints"`
	Estimates estimate.Sessions `json:"estimates"`
	Standups  standup.Sessions  `json:"standups"`
	Retros    retro.Sessions    `json:"retros"`
	Comments  comment.Comments  `json:"comments"`
}

func (er *EmailResponse) Subject() string {
	return fmt.Sprintf("[%v] %v report", npncore.ToYMD(er.Date), npncore.AppName)
}

func (er *EmailResponse) Opener() string {
	const msg = "Today there were %v teams, %v sprints, %v estimates, %v standups, %v retros, and %v comments"
	return fmt.Sprintf(msg, len(er.Teams), len(er.Sprints), len(er.Estimates), len(er.Standups), len(er.Retros), len(er.Comments))
}

var Email = Transcript{
	Key:         "email",
	Title:       "Email",
	Description: "Nightly email report",
	Resolve: func(ai npnweb.AppInfo, userID uuid.UUID, param string) (interface{}, error) {
		if len(param) == 0 {
			n := time.Now()
			param = npncore.ToYMD(&n)
		}
		d, err := npncore.FromYMD(param)
		if err != nil {
			return nil, err
		}
		svc := app.Svc(ai)
		return EmailResponse{
			Date:      d,
			Teams:     svc.Team.GetByCreated(d, nil),
			Sprints:   svc.Sprint.GetByCreated(d, nil),
			Estimates: svc.Estimate.GetByCreated(d, nil),
			Standups:  svc.Standup.GetByCreated(d, nil),
			Retros:    svc.Retro.GetByCreated(d, nil),
			Comments:  svc.Comment.GetByCreated(d, nil),
		}, nil
	},
}
