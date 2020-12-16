package recordstore

import (
	"context"
)

type Service interface {
	CreateRecord(c context.Context) (string, error)
	GetRecord(c context.Context) (string, error)
	SellRecord(c context.Context) (string, error)
	DeleteRecord(c context.Context) (string, error)
}
