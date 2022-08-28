package queryservice

import (
	"context"

	"github.com/s-beats/graphql-todo/domain"
)

type User interface {
	GetByID(ctx context.Context, id domain.UserID) (*domain.User, error)
}

type user struct{}

func NewUser() User {
	return &user{}
}

func (u *user) GetByID(ctx context.Context, id domain.UserID) (*domain.User, error) {
	return &domain.User{}, nil
}
