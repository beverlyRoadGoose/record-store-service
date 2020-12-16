package artistsmanager

import (
	"github.com/google/uuid"
	"heytobi.dev/record-store-service/internal"
)

func CreateArtist(name string) (internal.Artist, error) {
	return internal.Artist{
		Id:   uuid.New(),
		Name: name,
	}, nil
}
