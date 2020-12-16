package entities

import "github.com/google/uuid"

type Artist struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
