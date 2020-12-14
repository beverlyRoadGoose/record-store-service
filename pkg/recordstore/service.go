package recordstore

import (
	"context"
)

type Service interface {
	NewRecord(c context.Context)
	GetRecords(c context.Context)
	SellRecords(c context.Context)
	DeleteRecord(c context.Context)
}