package retro

import (
	"fmt"

	"github.com/kyleu/npn/npncore"
	"github.com/kyleu/npn/npndatabase"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/kyleu/rituals.dev/app/action"
	"github.com/kyleu/rituals.dev/app/util"
)

var defaultFeedbackOrdering = npncore.Orderings{{Column: npncore.KeyCategory, Asc: true}, {Column: npncore.KeyIdx, Asc: true}, {Column: npncore.KeyCreated, Asc: false}}

func (s *Service) GetFeedback(retroID uuid.UUID, params *npncore.Params) Feedbacks {
	params = npncore.ParamsWithDefaultOrdering(util.KeyFeedback, params, defaultFeedbackOrdering...)
	var dtos []feedbackDTO
	q := npndatabase.SQLSelect("*", util.KeyFeedback, "retro_id = $1", params.OrderByString(), params.Limit, params.Offset)
	err := s.db.Select(&dtos, q, nil, retroID)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error retrieving feedback for retro [%v]: %+v", retroID, err))
		return nil
	}
	return toFeedbacks(dtos)
}

func (s *Service) GetFeedbackByID(feedbackID uuid.UUID) (*Feedback, error) {
	dto := &feedbackDTO{}
	q := npndatabase.SQLSelectSimple("*", util.KeyFeedback, npncore.KeyID+" = $1")
	err := s.db.Get(dto, q, nil, feedbackID)
	if err != nil {
		return nil, err
	}
	return dto.toFeedback(), nil
}

func (s *Service) GetFeedbackRetroID(feedbackID uuid.UUID) (*uuid.UUID, error) {
	ret := uuid.UUID{}
	q := npndatabase.SQLSelectSimple(npncore.WithDBID(s.svc.Key), util.KeyFeedback, npncore.KeyID+" = $1")
	err := s.db.Get(&ret, q, nil, feedbackID)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Service) NewFeedback(retroID uuid.UUID, category string, content string, userID uuid.UUID) (*Feedback, error) {
	id := npncore.UUID()
	html := util.ToHTML(content)

	q := `insert into feedback (id, retro_id, idx, user_id, category, content, html) values (
    $1, $2, coalesce((select max(idx) + 1 from feedback p2 where p2.retro_id = $3 and p2.category = $4), 0), $5, $6, $7, $8
	)`
	err := s.db.Insert(q, nil, id, retroID, retroID, category, userID, category, content, html)
	if err != nil {
		return nil, err
	}

	actionContent := map[string]interface{}{"feedbackID": id}
	s.Data.Actions.Post(s.svc.Key, retroID, userID, action.ActFeedbackAdd, actionContent)

	return s.GetFeedbackByID(id)
}

func (s *Service) UpdateFeedback(feedbackID uuid.UUID, category string, content string, userID uuid.UUID) (*Feedback, error) {
	html := util.ToHTML(content)

	q := npndatabase.SQLUpdate(util.KeyFeedback, []string{npncore.KeyCategory, npncore.KeyContent, npncore.KeyHTML}, npncore.KeyID+" = $4")
	err := s.db.UpdateOne(q, nil, category, content, html, feedbackID)
	if err != nil {
		return nil, err
	}

	fb, err := s.GetFeedbackByID(feedbackID)
	if err != nil {
		return nil, errors.Wrap(err, "cannot load feedback ["+feedbackID.String()+"] for update")
	}
	if fb == nil {
		return nil, errors.New("cannot load newly-updated feedback")
	}

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.Data.Actions.Post(s.svc.Key, fb.RetroID, userID, action.ActFeedbackUpdate, actionContent)

	return s.GetFeedbackByID(feedbackID)
}

func (s *Service) RemoveFeedback(feedbackID uuid.UUID, userID uuid.UUID) error {
	feedback, err := s.GetFeedbackByID(feedbackID)
	if err != nil {
		return errors.Wrap(err, "cannot load feedback ["+feedbackID.String()+"] for removal")
	}
	if feedback == nil {
		return errors.New("cannot load feedback [" + feedbackID.String() + "] for removal")
	}

	q := npndatabase.SQLDelete(util.KeyFeedback, npncore.KeyID+" = $1")
	err = s.db.DeleteOne(q, nil, feedbackID)

	actionContent := map[string]interface{}{"feedbackID": feedbackID}
	s.Data.Actions.Post(s.svc.Key, feedback.RetroID, userID, action.ActFeedbackRemove, actionContent)

	return err
}

func toFeedbacks(dtos []feedbackDTO) Feedbacks {
	ret := make(Feedbacks, 0, len(dtos))
	for _, dto := range dtos {
		ret = append(ret, dto.toFeedback())
	}
	return ret
}
