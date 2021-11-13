package repository

import "github.com/s-beats/graphql-todo/domain/model"

type Task interface {
	Create(task *model.Task) (*model.Task, error)
}
