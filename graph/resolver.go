package graph

import "github.com/s-beats/graphql-todo/usecase"

type Resolver struct {
	TaskUsecase usecase.Task
}
