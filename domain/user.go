package domain

import "github.com/s-beats/graphql-todo/util"

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
