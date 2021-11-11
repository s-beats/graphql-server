package repository

import "github.com/s-beats/graphql-todo/graph/model"

type Task interface {
	Create(task *model.Task) (*model.Task, error)
}
