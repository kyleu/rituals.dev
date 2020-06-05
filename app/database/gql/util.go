package gql

import (
	"fmt"

	"github.com/kyleu/rituals.dev/app/util"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/kyleu/rituals.dev/app/database/query"
	"github.com/kyleu/rituals.dev/app/web"
	"logur.dev/logur"
)

var graphQLInitialized = false

type Callback func(p graphql.ResolveParams, ctx web.RequestContext) (interface{}, error)

func paramSetFromGraphQLParams(key string, params graphql.ResolveParams, logger logur.Logger) *query.Params {
	orderings := make(query.Orderings, 0)
	o, ok := params.Args["orders"]
	if ok {
		for _, x := range o.([]interface{}) {
			m := x.(map[string]interface{})
			col := util.MapGetString(m, "col", logger)
			asc := util.MapGetBool(m, "asc", logger)
			var defaultOrdering = query.Orderings{{Column: col, Asc: asc}}
			orderings = append(orderings, defaultOrdering...)
		}
	}

	limit := 0
	l, ok := params.Args["limit"]
	if ok {
		limit = l.(int)
	}

	offset := 0
	x, ok := params.Args["offset"]
	if ok {
		offset = x.(int)
	}

	ret := &query.Params{Key: key, Orderings: orderings, Limit: limit, Offset: offset}
	return ret.Filtered(logger)
}

func ErrorResponseJSON(logger logur.Logger, errors ...error) *graphql.Result {
	var errs = make([]gqlerrors.FormattedError, 0, len(errors))

	for _, err := range errors {
		logger.Warn(fmt.Sprintf("error running GraphQL: %+v", err))
		errs = append(errs, gqlerrors.FormattedError{Message: err.Error()})
	}

	return &graphql.Result{
		Errors: errs,
	}
}
