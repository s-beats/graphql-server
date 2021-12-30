package usecase

import (
	"context"

	"github.com/s-beats/graphql-todo/domain"
	"github.com/s-beats/graphql-todo/domain/repository"
	"github.com/s-beats/graphql-todo/util"
)

type Task interface {
	Create(ctx context.Context, title, text, userID, priority string) (*domain.Task, error)
}

type task struct {
	taskRepository repository.Task
}

func NewTask(taskRepo repository.Task) Task {
	return &task{
		taskRepository: taskRepo,
	}
}

func (t *task) Create(ctx context.Context, title, text, userID, priority string) (*domain.Task, error) {
	now := util.GetTimeNow()
	task := domain.NewTask(
		domain.NewTaskID(util.NewUUID()),
		domain.NewTaskTitle(title),
		domain.NewTaskText(text),
		now,
		now,
		domain.NewUser(domain.NewUserID(userID)),
		domain.NewPriority(priority),
	)

	err := t.taskRepository.Save(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}
