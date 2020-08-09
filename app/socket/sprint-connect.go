package socket

import (
	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/npn/npnuser"
	"github.com/kyleu/npn/npnservice/auth"
	"github.com/kyleu/rituals.dev/app/comment"
	"github.com/kyleu/rituals.dev/app/estimate"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/permission"
	"github.com/kyleu/rituals.dev/app/retro"
	"github.com/kyleu/rituals.dev/app/sprint"
	"github.com/kyleu/rituals.dev/app/standup"
	"github.com/kyleu/rituals.dev/app/team"
	"github.com/kyleu/rituals.dev/app/util"
)

type SprintSessionJoined struct {
	Profile     *npnuser.Profile          `json:"profile"`
	Session     *sprint.Session        `json:"session"`
	Comments    comment.Comments       `json:"comments"`
	Team        *team.Session          `json:"team"`
	Members     member.Entries         `json:"members"`
	Online      []uuid.UUID            `json:"online"`
	Estimates   estimate.Sessions      `json:"estimates"`
	Standups    standup.Sessions       `json:"standups"`
	Retros      retro.Sessions         `json:"retros"`
	Auths       auth.Displays          `json:"auths"`
	Permissions permission.Permissions `json:"permissions"`
}

func onSprintConnect(s *Service, conn *connection, sprintID uuid.UUID) error {
	ch := Channel{Svc: util.SvcSprint, ID: sprintID}
	err := s.Join(conn.ID, ch)
	if err != nil {
		return errors.Wrap(err, "error joining channel")
	}
	err = joinSprintSession(s, conn, ch)
	return errors.Wrap(err, "error joining sprint session")
}

func joinSprintSession(s *Service, conn *connection, ch Channel) error {
	dataSvc := s.sprints

	if ch.Svc != util.SvcSprint {
		return errors.New("sprint cannot handle [" + ch.Svc.Key + "] message")
	}

	sess := dataSvc.GetByID(ch.ID)
	if sess == nil {
		return errorNoSession(s, ch.Svc, conn.ID, ch.ID)
	}
	res := getSessionResult(s, sess.TeamID, nil, ch, conn)
	if res.Error != nil {
		return res.Error
	}

	sj := SprintSessionJoined{
		Profile:     &conn.Profile,
		Session:     sess,
		Comments:    dataSvc.Data.GetComments(ch.ID, nil),
		Team:        getTeamOpt(s, sess.TeamID),
		Members:     dataSvc.Data.Members.GetByModelID(ch.ID, nil),
		Online:      s.GetOnline(ch),
		Estimates:   s.estimates.GetBySprintID(ch.ID, nil),
		Standups:    s.standups.GetBySprintID(ch.ID, nil),
		Retros:      s.retros.GetBySprintID(ch.ID, nil),
		Auths:       res.Auth,
		Permissions: res.Perms,
	}
	msg := NewMessage(util.SvcSprint, ServerCmdSessionJoined, sj)
	return s.sendInitial(ch, conn, res.Entry, msg, nil, nil)
}
