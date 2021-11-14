package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/s-beats/graphql-todo/graph"
	"github.com/s-beats/graphql-todo/graph/generated"
	"github.com/s-beats/graphql-todo/usecase"
)

func GraphQLHandler(taskUsecase usecase.Task) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					TaskUsecase: taskUsecase,
				}},
		),
		)

		// extention handler
		// srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		// 	rawQuery := graphql.GetOperationContext(ctx).RawQuery

		// 	// format query
		// 	rp := regexp.MustCompile(`\n *| {2,}`)
		// 	q := rp.ReplaceAllString(rawQuery, " ")
		// 	// trim right space
		// 	q = strings.TrimRight(q, " ")
		// 	log.Debug().Str("query", q).Send()

		// 	return next(ctx)
		// })

		srv.ServeHTTP(w, r)
	}
}
