package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/s-beats/graphql-todo/domain"
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type Task interface {
	Save(ctx context.Context, task *domain.Task) error
	GetOne(ctx context.Context, taskID domain.TaskID) (*domain.Task, error)
}

type task struct {
	db *xorm.Engine
}

func NewTask(db *xorm.Engine) Task {
	return &task{
		db: db,
	}
}

type TaskDTO struct {
	ID         string `xorm:"'id'"`
	Title      string
	Text       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	UserID     string `xorm:"'user_id'"`
	PriorityID string `xorm:"'priority_id'"`
}

func (t *TaskDTO) TableName() string {
	return "tasks"
}

func (t *task) Save(ctx context.Context, task *domain.Task) error {
	taskDTO := &TaskDTO{
		ID:        task.ID().String(),
		Title:     task.Title().String(),
		Text:      task.Text().String(),
		CreatedAt: task.CreatedAt(),
		UpdatedAt: task.UpdatedAt(),
		UserID:    task.CreatedBy().ID().String(),
	}

	exists, err := t.db.Table(taskDTO.TableName()).Where("id = ?", taskDTO.ID).Exist()
	if err != nil {
		return xerrors.Errorf("%v", err)
	}

	var taskPriorityID int
	has, err := t.db.Table("task_priorities").Cols("id").Where("value = ?", task.Priority().Value()).Get(&taskPriorityID)
	if err != nil {
		return xerrors.Errorf("%v", err)
	}
	if !has {
		return xerrors.New("invalid task priority")
	}
	taskDTO.PriorityID = strconv.Itoa(taskPriorityID)

	if exists {
		_, err := t.db.ID(taskDTO.ID).Update(taskDTO)
		if err != nil {
			return xerrors.Errorf("%v", err)
		}
	} else {
		_, err := t.db.Insert(taskDTO)
		if err != nil {
			return xerrors.Errorf("%v", err)
		}
	}

	return nil
}

func (t *task) GetOne(ctx context.Context, taskID domain.TaskID) (*domain.Task, error) {
	panic("not implement")
}
