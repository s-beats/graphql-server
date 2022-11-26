package usecase

import (
	"context"

	"github.com/s-beats/graphql-todo/infra/rdb"
	"github.com/s-beats/graphql-todo/util"
)

type User interface {
	Create(ctx context.Context, name string) (*rdb.User, error)
}

type user struct {
	*rdb.Queries
}

func NewUser(q *rdb.Queries) User {
	return &user{
		Queries: q,
	}
}

func (u *user) Create(ctx context.Context, name string) (*rdb.User, error) {
	id := util.NewUUID()

	affected, err := u.Queries.CreateUser(ctx, rdb.CreateUserParams{
		ID:   id,
		Name: name,
	})
	if err != nil {
		return nil, err
	}
	if affected == 0 {
		return nil, errorFailedUpdateData()
	}

	user, err := u.Queries.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
