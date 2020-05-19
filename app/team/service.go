package team

import (
	"database/sql"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/query"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/kyleu/rituals.dev/app/member"
	"github.com/kyleu/rituals.dev/app/util"
	"logur.dev/logur"
)

type Service struct {
	actions *action.Service
	db      *sqlx.DB
	Members *member.Service
	logger  logur.Logger
}

func NewService(actions *action.Service, db *sqlx.DB, logger logur.Logger) *Service {
	logger = logur.WithFields(logger, map[string]interface{}{"service": util.SvcTeam.Key})

	return &Service{
		actions: actions,
		db:      db,
		Members: member.NewService(actions, db, util.SvcTeam.Key),
		logger:  logger,
	}
}

func (s *Service) New(title string, userID uuid.UUID) (*Session, error) {
	slug, err := member.NewSlugFor(s.db, util.SvcTeam.Key, title)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error creating team slug"))
	}

	model := NewSession(title, slug, userID)

	q := "insert into team (id, slug, title, owner) values ($1, $2, $3, $4)"
	_, err = s.db.Exec(q, model.ID, slug, model.Title, model.Owner)
	if err != nil {
		return nil, errors.WithStack(errors.Wrap(err, "error saving new team session"))
	}

	s.actions.Post(util.SvcTeam.Key, model.ID, userID, action.ActCreate, nil, "")
	return &model, nil
}

func (s *Service) List(params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcTeam.Key, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.SvcTeam.Key, "", params.OrderByString(), params.Limit, params.Offset))
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByID(id uuid.UUID) (*Session, error) {
	dto := &sessionDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.SvcTeam.Key, "id = $1", "", 0, 0), id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetBySlug(slug string) (*Session, error) {
	var dto = &sessionDTO{}
	err := s.db.Get(dto, query.SQLSelect("*", util.SvcTeam.Key, "slug = $1", "", 0, 0), slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return dto.ToSession(), nil
}

func (s *Service) GetByOwner(userID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcTeam.Key, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	err := s.db.Select(&dtos, query.SQLSelect("*", util.SvcTeam.Key, "owner = $1", params.OrderByString(), params.Limit, params.Offset), userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetByMember(userID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcTeam.Key, params, &query.Ordering{Column: "m.created", Asc: false})
	var dtos []sessionDTO
	q := query.SQLSelect("x.*", "team x join team_member m on x.id = m.team_id", "m.user_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, userID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) GetBySprint(sprintID uuid.UUID, params *query.Params) ([]*Session, error) {
	params = query.ParamsWithDefaultOrdering(util.SvcTeam.Key, params, &query.Ordering{Column: "created", Asc: false})
	var dtos []sessionDTO
	q := query.SQLSelect("*", util.SvcTeam.Key, "sprint_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, sprintID)
	if err != nil {
		return nil, err
	}
	return toSessions(dtos), nil
}

func (s *Service) UpdateSession(sessionID uuid.UUID, title string, userID uuid.UUID) error {
	q := "update team set title = $1 where id = $2"
	_, err := s.db.Exec(q, title, sessionID)
	s.actions.Post(util.SvcTeam.Key, sessionID, userID, action.ActUpdate, nil, "")
	return errors.WithStack(errors.Wrap(err, "error updating team session"))
}

func toSessions(dtos []sessionDTO) []*Session {
	ret := make([]*Session, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.ToSession())
	}
	return ret
}
