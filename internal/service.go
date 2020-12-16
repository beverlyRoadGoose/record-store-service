package internal

import (
	"context"
	"github.com/google/uuid"
	"heytobi.dev/record-store-service/internal/entities"
)

type Service interface {
	CreateArtist(c context.Context, name string) (entities.Artist, error)
	GetArtist(c context.Context) (string, error)
	DeleteArtist(c context.Context) (string, error)
	CreateRecord(c context.Context, aid uuid.UUID, name string, year uint) (entities.Record, error)
	GetRecord(c context.Context) (string, error)
	SellRecord(c context.Context) (string, error)
	DeleteRecord(c context.Context) (string, error)
}
