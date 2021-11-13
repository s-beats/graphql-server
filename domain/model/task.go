package model

import (
	"time"

	"github.com/s-beats/graphql-todo/util"
)

type Task struct {
	ID        TaskID
	Title     TaskTitle
	Text      TaskText
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(title, text string) *Task {
	now := util.GetTimeNow()
	return &Task{
		ID:        *NewTaskID(),
		Title:     *NewTaskTitle(title),
		Text:      *NewTaskText(text),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

type TaskID struct {
	id string
}

func NewTaskID() *TaskID {
	return &TaskID{
		id: util.NewUUID(),
	}
}

type TaskTitle struct {
	title string
}

func NewTaskTitle(title string) *TaskTitle {
	return &TaskTitle{
		title: title,
	}
}

type TaskText struct {
	text string
}

func NewTaskText(text string) *TaskText {
	return &TaskText{
		text: text,
	}
}
