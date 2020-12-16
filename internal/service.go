package internal

import (
	"context"
	"heytobi.dev/record-store-service/internal/entities"
)

type Service interface {
	CreateArtist(c context.Context, name string) (entities.Artist, error)
	GetArtist(c context.Context) (string, error)
	DeleteArtist(c context.Context) (string, error)
	CreateRecord(c context.Context) (string, error)
	GetRecord(c context.Context) (string, error)
	SellRecord(c context.Context) (string, error)
	DeleteRecord(c context.Context) (string, error)
}
