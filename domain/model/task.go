package model

import "time"

type Task struct {
	ID        TaskID
	Title     TaskTitle
	Text      TaskText
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TaskID struct {
	id string
}

type TaskTitle struct {
	title string
}

type TaskText struct {
	title string
}
