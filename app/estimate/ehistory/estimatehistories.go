// Package ehistory - Content managed by Project Forge, see [projectforge.md] for details.
package ehistory

import (
	"slices"

	"github.com/google/uuid"
	"github.com/samber/lo"
)

type EstimateHistories []*EstimateHistory

func (e EstimateHistories) Get(slug string) *EstimateHistory {
	return lo.FindOrElse(e, nil, func(x *EstimateHistory) bool {
		return x.Slug == slug
	})
}

func (e EstimateHistories) GetBySlugs(slugs ...string) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return lo.Contains(slugs, xx.Slug)
	})
}

func (e EstimateHistories) GetBySlug(slug string) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return xx.Slug == slug
	})
}

func (e EstimateHistories) GetByEstimateIDs(estimateIDs ...uuid.UUID) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return lo.Contains(estimateIDs, xx.EstimateID)
	})
}

func (e EstimateHistories) GetByEstimateID(estimateID uuid.UUID) EstimateHistories {
	return lo.Filter(e, func(xx *EstimateHistory, _ int) bool {
		return xx.EstimateID == estimateID
	})
}

func (e EstimateHistories) Slugs() []string {
	return lo.Map(e, func(x *EstimateHistory, _ int) string {
		return x.Slug
	})
}

func (e EstimateHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(e)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(e, func(x *EstimateHistory, _ int) {
		ret = append(ret, x.Slug)
	})
	return ret
}

func (e EstimateHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(e)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(e, func(x *EstimateHistory, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (e EstimateHistories) Clone() EstimateHistories {
	return slices.Clone(e)
}
