package recordstore

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
)

type recordStoreService struct{}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

func NewRecordStoreService() Service {
	return &recordStoreService{}
}

/*
 *
 */
func (r recordStoreService) NewRecord(_ context.Context) {
	panic("implement me")
}

/*
 *
 */
func (r recordStoreService) GetRecords(_ context.Context) {
	panic("implement me")
}

/*
 *
 */
func (r recordStoreService) SellRecords(_ context.Context) {
	panic("implement me")
}

/*
 *
 */
func (r recordStoreService) DeleteRecord(_ context.Context) {
	panic("implement me")
}
