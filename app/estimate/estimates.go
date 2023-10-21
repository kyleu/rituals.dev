// Package estimate - Content managed by Project Forge, see [projectforge.md] for details.
package estimate

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/rituals/app/enum"
)

type Estimates []*Estimate

func (e Estimates) Get(id uuid.UUID) *Estimate {
	return lo.FindOrElse(e, nil, func(x *Estimate) bool {
		return x.ID == id
	})
}

func (e Estimates) GetByIDs(ids ...uuid.UUID) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return lo.Contains(ids, xx.ID)
	})
}

func (e Estimates) GetByID(id uuid.UUID) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return xx.ID == id
	})
}

func (e Estimates) GetBySlugs(slugs ...string) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (e Estimates) GetBySlug(slug string) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return xx.Slug == slug
	})
}

func (e Estimates) GetByStatuses(statuses ...enum.SessionStatus) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return lo.Contains(statuses, xx.Status)
	})
}

func (e Estimates) GetByStatus(status enum.SessionStatus) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return xx.Status == status
	})
}

func (e Estimates) GetByTeamIDs(teamIDs ...*uuid.UUID) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return lo.Contains(teamIDs, xx.TeamID)
	})
}

func (e Estimates) GetByTeamID(teamID *uuid.UUID) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return xx.TeamID == teamID
	})
}

func (e Estimates) GetBySprintIDs(sprintIDs ...*uuid.UUID) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return lo.Contains(sprintIDs, xx.SprintID)
	})
}

func (e Estimates) GetBySprintID(sprintID *uuid.UUID) Estimates {
	return lo.Filter(e, func(xx *Estimate, _ int) bool {
		return xx.SprintID == sprintID
	})
}

func (e Estimates) IDs() []uuid.UUID {
	return lo.Map(e, func(x *Estimate, _ int) uuid.UUID {
		return x.ID
	})
}

func (e Estimates) IDStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *Estimate, _ int) {
		ret = append(ret, x.ID.String())
	})
	return ret
}

func (e Estimates) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(e, func(x *Estimate, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (e Estimates) Clone() Estimates {
	return slices.Clone(e)
}
