package queryservice

import (
	"context"
	"errors"

	"github.com/s-beats/graphql-todo/domain"
)

var dummyTasks = map[domain.TaskID]*domain.Task{
	*domain.NewTaskID("3b671140-c499-4750-82f5-f015faae25e5"): &domain.Task{},
	*domain.NewTaskID("c49c0c71-2739-4979-bd6c-939aad5bf6b0"): &domain.Task{},
}

type Task interface {
	GetByID(ctx context.Context, id domain.TaskID) (*domain.Task, error)
}

type task struct{}

func NewTask() Task {
	return &task{}
}

func (t *task) GetByID(ctx context.Context, id domain.TaskID) (*domain.Task, error) {
	task, ok := dummyTasks[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return task, nil
}
