package graph

import (
	"github.com/s-beats/graphql-todo/queryservice"
	"github.com/s-beats/graphql-todo/usecase"
)

type Resolver struct {
	TaskUsecase usecase.Task
	UserUsecase usecase.User

	UserQueryService queryservice.User
	TaskQueryService queryservice.Task
}
