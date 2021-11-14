package internal

import (
	"github.com/s-beats/graphql-todo/domain"
	"github.com/s-beats/graphql-todo/graph/model"
)

func ConvertTask(task *domain.Task) *model.Task {
	return &model.Task{
		ID:        task.ID().String(),
		Title:     task.Title().String(),
		Text:      task.Text().String(),
		CreatedAt: task.CreatedAt(),
		UpdatedAt: task.UpdatedAt(),
	}
}
