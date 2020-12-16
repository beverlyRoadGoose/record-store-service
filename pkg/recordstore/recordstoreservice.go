package recordstore

import (
	"context"
)

type recordStoreService struct{}

func NewRecordStoreService() Service {
	return &recordStoreService{}
}

func (r recordStoreService) CreateRecord(_ context.Context) (string, error) {
	return "Created new record.", nil
}

func (r recordStoreService) GetRecord(_ context.Context) (string, error) {
	return "Record.", nil
}

func (r recordStoreService) SellRecord(_ context.Context) (string, error) {
	return "Record sold.", nil
}

func (r recordStoreService) DeleteRecord(_ context.Context) (string, error) {
	return "Record deleted.", nil
}
