package domain

import (
	"github.com/s-beats/graphql-todo/util"
)

type User struct {
	id   *UserID
	name string
}

func NewUser(id *UserID, name string) *User {
	return &User{
		id:   id,
		name: name,
	}
}

func (u *User) ID() *UserID {
	return u.id
}

type UserID struct {
	id string
}

func (u *UserID) String() string {
	return u.id
}

func NewUserID(id string) *UserID {
	return &UserID{
		id: util.UUIDMustParse(id),
	}
}

func (u *User) Name() string {
	return u.name
}
