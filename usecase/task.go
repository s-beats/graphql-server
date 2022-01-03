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
	userRepository repository.User
}

func NewTask(taskRepo repository.Task, userRepo repository.User) Task {
	return &task{
		taskRepository: taskRepo,
		userRepository: userRepo,
	}
}

func (t *task) Create(ctx context.Context, title, text, userID, priority string) (*domain.Task, error) {
	user, err := t.userRepository.GetOne(ctx, *domain.NewUserID(userID))
	if err != nil {
		return nil, err
	}

	now := util.GetTimeNow()
	task := domain.NewTask(
		domain.NewTaskID(util.NewUUID()),
		domain.NewTaskTitle(title),
		domain.NewTaskText(text),
		now,
		now,
		user,
		domain.NewPriority(priority),
	)

	if err := t.taskRepository.Save(ctx, task); err != nil {
		return nil, err
	}

	return task, nil
}
