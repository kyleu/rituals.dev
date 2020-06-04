package util

import "strings"

type Service struct {
	Key         string
	Title       string
	Plural      string
	PluralTitle string
	Description string
	Icon        string
}

// Services

var SvcSystem = Service{
	Key:         KeySystem,
	Title:       Title(KeySystem),
	Plural:      Plural(KeySystem),
	PluralTitle: PluralTitle(KeySystem),
	Description: "",
	Icon:        "close",
}
var SvcTeam = Service{
	Key:         "team",
	Title:       "Team",
	Plural:      "teams",
	PluralTitle: "Teams",
	Description: "Join your friends and work towards a common goal",
	Icon:        "users",
}
var SvcSprint = Service{
	Key:         "sprint",
	Title:       "Sprint",
	Plural:      "sprints",
	PluralTitle: "Sprints",
	Description: "Plan your time and direct your efforts",
	Icon:        "settings",
}
var SvcEstimate = Service{
	Key:         "estimate",
	Title:       "Estimate",
	Plural:      "estimates",
	PluralTitle: "Estimates",
	Description: "Planning poker for any stories you need to work on",
	Icon:        "tag",
}
var SvcStandup = Service{
	Key:         "standup",
	Title:       "Standup",
	Plural:      "standups",
	PluralTitle: "Standups",
	Description: "Share your progress with your team",
	Icon:        "future",
}
var SvcRetro = Service{
	Key:         "retro",
	Title:       "Retrospective",
	Plural:      "retros",
	PluralTitle: "Retrospectives",
	Description: "Discover improvements and praise for your work",
	Icon:        "history",
}

var allServices = []Service{SvcSystem, SvcTeam, SvcSprint, SvcEstimate, SvcStandup, SvcRetro}

func ServiceFromString(str string) Service {
	for _, s := range allServices {
		if s.Key == str {
			return s
		}
	}
	return SvcSystem
}

func ServiceTitle(svc Service, title string) string {
	title = strings.TrimSpace(title)
	if len(title) == 0 {
		title = "Untitled " + svc.Title
	}

	return title
}
