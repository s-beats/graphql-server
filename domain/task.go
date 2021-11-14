package domain

import (
	"time"

	"github.com/s-beats/graphql-todo/util"
)

type Task struct {
	id        *TaskID
	title     *TaskTitle
	text      *TaskText
	createdAt time.Time
	updatedAt time.Time
	createdBy *UserID
}

func NewTask(id *TaskID, title *TaskTitle, text *TaskText, createdAt, updatedAt time.Time, createdBy *UserID) *Task {
	return &Task{
		id:        id,
		title:     title,
		text:      text,
		createdAt: createdAt,
		updatedAt: updatedAt,
		createdBy: createdBy,
	}
}

func (t *Task) ID() *TaskID {
	return t.id
}

func (t *Task) Title() *TaskTitle {
	return t.title
}

func (t *Task) Text() *TaskText {
	return t.text
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) CreatedBy() *UserID {
	return t.createdBy
}

type TaskID struct {
	id string
}

func (t *TaskID) String() string {
	return t.id
}

func NewTaskID(id string) *TaskID {
	return &TaskID{
		id: util.UUIDMustParse(id),
	}
}

type TaskTitle struct {
	title string
}

func (t *TaskTitle) String() string {
	return t.title
}

func NewTaskTitle(title string) *TaskTitle {
	return &TaskTitle{
		title: title,
	}
}

type TaskText struct {
	text string
}

func (t *TaskText) String() string {
	return t.text
}

func NewTaskText(text string) *TaskText {
	return &TaskText{
		text: text,
	}
}
