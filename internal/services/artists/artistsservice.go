package artists

import (
	"github.com/google/uuid"
	"heytobi.dev/record-store-service/internal/entities"
)

func Create(name string) (entities.Artist, error) {
	return entities.Artist{
		Id:   uuid.New(),
		Name: name,
	}, nil
}
