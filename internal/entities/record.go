package entities

import (
	"github.com/google/uuid"
)

type Record struct {
	ArtistId uuid.UUID `json:"artistId"`
	Name     string    `json:"name"`
	Year     uint      `json:"year"`
}
