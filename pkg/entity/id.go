package entity

import (
	"github.com/google/uuid"
)

func NewId() uuid.UUID {
	return uuid.New()
}

func ParseID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)

	return id, err
}
