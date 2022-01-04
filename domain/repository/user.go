package repository

import (
	"context"
	"time"

	"github.com/s-beats/graphql-todo/domain"
	"golang.org/x/xerrors"
	"xorm.io/xorm"
)

type User interface {
	Save(ctx context.Context, user *domain.User) error
	GetOne(ctx context.Context, userID domain.UserID) (*domain.User, error)
}

type user struct {
	db *xorm.Engine
}

func NewUser(db *xorm.Engine) User {
	return &user{
		db: db,
	}
}

type UserDTO struct {
	ID        string `xorm:"'id'"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *UserDTO) TableName() string {
	return "users"
}

func (t *user) Save(ctx context.Context, user *domain.User) error {
	userDTO := &UserDTO{
		ID:        user.ID().String(),
		Name:      user.Name(),
		UpdatedAt: time.Now(),
	}

	exists, err := t.db.Table(userDTO.TableName()).Where("id = ?", userDTO.ID).Exist()
	if err != nil {
		return xerrors.Errorf("%v", err)
	}

	if exists {
		_, err := t.db.ID(userDTO.ID).Update(userDTO)
		if err != nil {
			return xerrors.Errorf("%v", err)
		}
	} else {
		userDTO.CreatedAt = userDTO.UpdatedAt
		_, err := t.db.Insert(userDTO)
		if err != nil {
			return xerrors.Errorf("%v", err)
		}
	}

	return nil
}

func (t *user) GetOne(ctx context.Context, userID domain.UserID) (*domain.User, error) {
	userDTO := UserDTO{ID: userID.String()}
	has, err := t.db.Get(&userDTO)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, xerrors.New("user not found")
	}
	return domain.NewUser(
		domain.NewUserID(userDTO.ID),
		userDTO.Name,
	), nil
}
