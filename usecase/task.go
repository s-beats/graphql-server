package usecase

import (
	"context"

	"github.com/s-beats/graphql-todo/infra/rdb"
	"github.com/s-beats/graphql-todo/util"
)

type Task interface {
	Create(ctx context.Context, title, text, userID, priority string) (*rdb.Task, error)
}

type task struct {
	*rdb.Queries
}

func NewTask(q *rdb.Queries) Task {
	return &task{
		Queries: q,
	}
}

func (t *task) Create(ctx context.Context, title, text, userID, priority string) (*rdb.Task, error) {
	user, err := t.Queries.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	taskPriority, err := t.Queries.GetTaskPriority(ctx, priority)
	if err != nil {
		return nil, err
	}

	id := util.NewUUID()
	affected, err := t.Queries.CreateTask(ctx, rdb.CreateTaskParams{
		ID:         id,
		Title:      title,
		Text:       text,
		UserID:     user.ID,
		PriorityID: taskPriority.ID,
	})
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errorFailedUpdateData()
	}

	task, err := t.Queries.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
