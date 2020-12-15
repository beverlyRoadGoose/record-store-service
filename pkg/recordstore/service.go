package recordstore

import (
	"context"
)

type Service interface {
	CreateRecords(c context.Context) (string, error)
	GetRecords(c context.Context) (string, error)
	SellRecords(c context.Context) (string, error)
	DeleteRecords(c context.Context) (string, error)
}
