package usecase

import (
	"context"

	"github.com/s-beats/graphql-todo/domain"
	"github.com/s-beats/graphql-todo/domain/repository"
	"github.com/s-beats/graphql-todo/util"
)

type User interface {
	Create(ctx context.Context, name string) (*domain.User, error)
}

type user struct {
	userRepository repository.User
}

func NewUser(userRepo repository.User) User {
	return &user{
		userRepository: userRepo,
	}
}

func (u *user) Create(ctx context.Context, name string) (*domain.User, error) {
	user := domain.NewUser(
		domain.NewUserID(util.NewUUID()),
		name,
	)

	err := u.userRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
