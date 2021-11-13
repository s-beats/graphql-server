package util

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.NewString()
}

func UUIDMustParse(id string) string {
	return uuid.MustParse(id).String()
}
