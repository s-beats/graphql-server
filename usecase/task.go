package usecase

import (
	"github.com/s-beats/graphql-todo/domain/model"
	"github.com/s-beats/graphql-todo/domain/repository"
)

type Task interface {
	Create(input *CreateTaskInput) (*CreateTaskOutput, error)
}

type task struct {
	taskRepository repository.Task
}

func NewTask(taskRepo repository.Task) Task {
	return &task{
		taskRepository: taskRepo,
	}
}

type CreateTaskInput struct {
	Title string
	Text  string
}

type CreateTaskOutput struct {
	Task *model.Task
}

func (t *task) Create(input *CreateTaskInput) (*CreateTaskOutput, error) {
	task := model.NewTask(input.Title, input.Text)

	createdTask, err := t.taskRepository.Create(task)
	if err != nil {
		return nil, err
	}

	return &CreateTaskOutput{
		Task: createdTask,
	}, nil
}
