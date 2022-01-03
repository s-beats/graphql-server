package internal

import (
	"github.com/s-beats/graphql-todo/domain"
	"github.com/s-beats/graphql-todo/graph/model"
)

func ConvertTask(task *domain.Task) *model.Task {
	priority := model.TaskPriority(task.Priority())
	if !priority.IsValid() {
		panic("can not convert, invalid priority value")
	}
	return &model.Task{
		ID:        task.ID().String(),
		Title:     task.Title().String(),
		Text:      task.Text().String(),
		CreatedAt: task.CreatedAt(),
		UpdatedAt: task.UpdatedAt(),
		Priority:  priority,
	}
}

func ConvertUser(user *domain.User) *model.User {
	return &model.User{
		ID:   user.ID().String(),
		Name: user.Name(),
	}
}
