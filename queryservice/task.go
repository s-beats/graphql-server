package queryservice

import (
	"context"

	"github.com/s-beats/graphql-todo/infra/rdb"
	"github.com/samber/lo"
)

type Task interface {
	Search(ctx context.Context, params rdb.SearchTasksParams) ([]*rdb.Task, error)
	GetByID(ctx context.Context, id string) (*rdb.Task, error)
}

type task struct {
	*rdb.Queries
}

func NewTask(q *rdb.Queries) Task {
	return &task{
		Queries: q,
	}
}

func (t *task) Search(ctx context.Context, params rdb.SearchTasksParams) ([]*rdb.Task, error) {
	tasks, err := t.Queries.SearchTasks(ctx, params)
	if err != nil {
		return nil, err
	}
	return lo.ToSlicePtr(tasks), nil
}

func (t *task) GetByID(ctx context.Context, id string) (*rdb.Task, error) {
	task, err := t.Queries.GetTask(ctx, id)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
