package util

import "github.com/google/uuid"

func NewUUID() string {
	return uuid.NewString()
}

func UUIDParse(id string) (string, error) {
	parsedID, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}
	return parsedID.String(), nil
}
