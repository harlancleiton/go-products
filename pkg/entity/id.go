package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func NewIDFromString(id string) (ID, error) {
	return uuid.Parse(id)
}
