// Content managed by Project Forge, see [projectforge.md] for details.
package thistory

import (
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
)

type TeamHistories []*TeamHistory

func (t TeamHistories) Get(slug string) *TeamHistory {
	return lo.FindOrElse(t, nil, func(x *TeamHistory) bool {
		return x.Slug == slug
	})
}

func (t TeamHistories) GetBySlugs(slugs ...string) TeamHistories {
	return lo.Filter(t, func(x *TeamHistory, _ int) bool {
		return lo.Contains(slugs, x.Slug)
	})
}

func (t TeamHistories) Slugs() []string {
	return lo.Map(t, func(x *TeamHistory, _ int) string {
		return x.Slug
	})
}

func (t TeamHistories) SlugStrings(includeNil bool) []string {
	ret := make([]string, 0, len(t)+1)
	if includeNil {
		ret = append(ret, "")
	}
	lo.ForEach(t, func(x *TeamHistory, _ int) {
		ret = append(ret, x.Slug)
	})
	return ret
}

func (t TeamHistories) TitleStrings(nilTitle string) []string {
	ret := make([]string, 0, len(t)+1)
	if nilTitle != "" {
		ret = append(ret, nilTitle)
	}
	lo.ForEach(t, func(x *TeamHistory, _ int) {
		ret = append(ret, x.TitleString())
	})
	return ret
}

func (t TeamHistories) Clone() TeamHistories {
	return slices.Clone(t)
}
