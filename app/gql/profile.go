package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/kyleu/rituals.dev/app/util"
	"github.com/kyleu/rituals.dev/app/web"
)

var (
	profileResolver         Callback
	profileSprintResolver   Callback
	profileEstimateResolver Callback
	profileStandupResolver  Callback
	profileRetroResolver    Callback
	profileType             *graphql.Object
)

func initProfile() {
	profileResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.Profile.ToProfile(), nil
	}

	profileSprintResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Sprint.GetByMember(p.Source.(util.Profile).UserID, paramSetFromGraphQLParams(util.SvcSprint.Key, p, ctx.Logger))
	}

	profileEstimateResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Estimate.GetByMember(p.Source.(util.Profile).UserID, paramSetFromGraphQLParams(util.SvcEstimate.Key, p, ctx.Logger))
	}

	profileStandupResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Standup.GetByMember(p.Source.(util.Profile).UserID, paramSetFromGraphQLParams(util.SvcStandup.Key, p, ctx.Logger))
	}

	profileRetroResolver = func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error) {
		return ctx.App.Retro.GetByMember(p.Source.(util.Profile).UserID, paramSetFromGraphQLParams(util.SvcRetro.Key, p, ctx.Logger))
	}

	profileType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Profile",
			Fields: graphql.Fields{
				"userID": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"role": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"theme": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"navColor": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"linkColor": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				"picture": &graphql.Field{
					Type: graphql.String,
				},
				"locale": &graphql.Field{
					Type: graphql.NewNonNull(graphql.String),
				},
				util.SvcSprint.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(sprintType)),
					Description: "Your current sprints",
					Args: listArgs,
					Resolve:     ctxF(profileSprintResolver),
				},
				util.SvcEstimate.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(estimateType)),
					Description: "Your current estimates",
					Args: listArgs,
					Resolve:     ctxF(profileEstimateResolver),
				},
				util.SvcStandup.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(standupType)),
					Description: "Your current standups",
					Args: listArgs,
					Resolve:     ctxF(profileStandupResolver),
				},
				util.SvcRetro.Plural: &graphql.Field{
					Type:        graphql.NewList(graphql.NewNonNull(retroType)),
					Description: "Your current retros",
					Args: listArgs,
					Resolve:     ctxF(profileRetroResolver),
				},
			},
		},
	)

	memberType.AddFieldConfig(util.KeyUser, &graphql.Field{
		Type:        profileType,
		Description: "The user associated to this member",
		Resolve:     ctxF(memberProfileResolver),
	})
}
