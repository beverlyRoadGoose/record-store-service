package recordstore

import (
	"context"
)

type recordStoreService struct{}

func NewRecordStoreService() Service {
	return &recordStoreService{}
}

func (r recordStoreService) CreateRecords(_ context.Context) (string, error) {
	return "Created new record.", nil
}

func (r recordStoreService) GetRecords(_ context.Context) (string, error) {
	return "Records.", nil
}

func (r recordStoreService) SellRecords(_ context.Context) (string, error) {
	return "Records sold.", nil
}

func (r recordStoreService) DeleteRecords(_ context.Context) (string, error) {
	return "Records deleted.", nil
}
