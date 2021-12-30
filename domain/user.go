package domain

import "github.com/s-beats/graphql-todo/util"

type User struct {
	id *UserID
}

func NewUser(id *UserID) *User {
	return &User{
		id: id,
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
