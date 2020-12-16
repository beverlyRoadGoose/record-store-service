package internal

import (
	"context"
	"github.com/google/uuid"
	"heytobi.dev/record-store-service/internal/entities"
	"heytobi.dev/record-store-service/internal/services/artists"
	"heytobi.dev/record-store-service/internal/services/records"
)

type recordStoreService struct{}

func NewRecordStoreService() Service {
	return &recordStoreService{}
}

func (r *recordStoreService) CreateArtist(c context.Context, name string) (entities.Artist, error) {
	artist, err := artists.Create(name)
	if err != nil {
		return entities.Artist{}, err
	}
	return artist, nil
}

func (r *recordStoreService) GetArtist(c context.Context) (string, error) {
	return "Artist.", nil
}

func (r *recordStoreService) DeleteArtist(c context.Context) (string, error) {
	return "Deleted Artist.", nil
}

func (r *recordStoreService) CreateRecord(_ context.Context, aid uuid.UUID, name string, year uint) (entities.Record, error) {
	record, err := records.Create(aid, name, year)
	if err != nil {
		return entities.Record{}, err
	}
	return record, nil
}

func (r *recordStoreService) GetRecord(_ context.Context) (string, error) {
	return "Record.", nil
}

func (r *recordStoreService) SellRecord(_ context.Context) (string, error) {
	return "Record sold.", nil
}

func (r *recordStoreService) DeleteRecord(_ context.Context) (string, error) {
	return "Record deleted.", nil
}
