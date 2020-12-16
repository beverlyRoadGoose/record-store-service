package records

import (
	"github.com/google/uuid"
	"heytobi.dev/record-store-service/internal/entities"
)

func Create(aid uuid.UUID, name string, year uint) (entities.Record, error) {
	return entities.Record{
		ArtistId: aid,
		Name:     name,
		Year:     year,
	}, nil
}
