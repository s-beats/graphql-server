package repository

import (
	"context"
	"time"

	"github.com/s-beats/graphql-todo/domain"
	"github.com/uptrace/bun"
)

type Task interface {
	Save(ctx context.Context, task *domain.Task) error
	GetOne(ctx context.Context, taskID domain.TaskID) (*domain.Task, error)
}

type task struct {
	db *bun.DB
}

func NewTask(db *bun.DB) Task {
	return &task{
		db: db,
	}
}

type TaskDTO struct {
	bun.BaseModel `bun:"tasks"`

	ID        string
	Title     string
	Text      string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
}

func (t *task) Save(ctx context.Context, task *domain.Task) error {
	taskDTO := TaskDTO{
		ID:        task.ID().String(),
		Title:     task.Title().String(),
		Text:      task.Text().String(),
		CreatedAt: task.CreatedAt(),
		UpdatedAt: task.UpdatedAt(),
		UserID:    task.CreatedBy().String(),
	}

	exists, err := t.db.NewSelect().Model((*TaskDTO)(nil)).Where("id = ?", taskDTO.ID).Exists(ctx)
	if err != nil {
		panic(err)
	}
	if exists {
		_, err := t.db.NewUpdate().
			Model(&taskDTO).
			Where("id = ?", taskDTO.ID).
			Exec(ctx)
		if err != nil {
			return err
		}
	} else {
		_, err := t.db.NewInsert().Model(&taskDTO).Exec(ctx)
		if err != nil {
			return err
		}

	}

	return nil
}

func (t *task) GetOne(ctx context.Context, taskID domain.TaskID) (*domain.Task, error) {
	panic("not implement")
}
