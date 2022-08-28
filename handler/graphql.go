package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/s-beats/graphql-todo/graph"
	"github.com/s-beats/graphql-todo/graph/generated"
	"github.com/s-beats/graphql-todo/queryservice"
	"github.com/s-beats/graphql-todo/usecase"
)

func GraphQLHandler(taskUsecase usecase.Task, userUsecase usecase.User, taskQueryService queryservice.Task, userQueryService queryservice.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graph.Resolver{
					TaskUsecase: taskUsecase,
					UserUsecase: userUsecase,
				}},
		),
		)

		srv.ServeHTTP(w, r)
	}
}
