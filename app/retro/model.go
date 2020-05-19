package retro

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/kyleu/rituals.dev/app/query"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/gofrs/uuid"
)

type Status struct {
	Key string `json:"key"`
}

var StatusNew = Status{Key: "new"}
var StatusDeleted = Status{Key: "deleted"}

var AllStatuses = []Status{StatusNew, StatusDeleted}

func statusFromString(s string) Status {
	for _, t := range AllStatuses {
		if t.Key == s {
			return t
		}
	}
	return StatusNew
}

func (t *Status) String() string {
	return t.Key
}

func (t Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Key)
}

var DefaultCategories = []string{"good", "bad", "improve"}

func categoriesFromDB(s string) []string {
	ret := query.StringToArray(s)
	if len(ret) == 0 {
		return DefaultCategories
	}
	return ret
}

type Session struct {
	ID         uuid.UUID  `json:"id"`
	Slug       string     `json:"slug"`
	Title      string     `json:"title"`
	SprintID   *uuid.UUID `json:"sprintID"`
	Owner      uuid.UUID  `json:"owner"`
	Status     Status     `json:"status"`
	Categories []string   `json:"categories"`
	Created    time.Time  `json:"created"`
}

func (s *Session) GetSlug() string {
	return s.Slug
}
func (s *Session) GetTitle() string {
	return s.Title
}


func NewSession(title string, slug string, userID uuid.UUID, sprintID *uuid.UUID) Session {
	return Session{
		ID:         util.UUID(),
		Slug:       slug,
		Title:      strings.TrimSpace(title),
		SprintID:   sprintID,
		Owner:      userID,
		Status:     StatusNew,
		Categories: make([]string, 0),
		Created:    time.Time{},
	}
}

type sessionDTO struct {
	ID         uuid.UUID  `db:"id"`
	Slug       string     `db:"slug"`
	Title      string     `db:"title"`
	SprintID   *uuid.UUID `db:"sprint_id"`
	Owner      uuid.UUID  `db:"owner"`
	Status     string     `db:"status"`
	Categories string     `db:"categories"`
	Created    time.Time  `db:"created"`
}

func (dto *sessionDTO) ToSession() *Session {
	return &Session{
		ID:         dto.ID,
		Slug:       dto.Slug,
		Title:      dto.Title,
		SprintID:   dto.SprintID,
		Owner:      dto.Owner,
		Status:     statusFromString(dto.Status),
		Categories: categoriesFromDB(dto.Categories),
		Created:    dto.Created,
	}
}
