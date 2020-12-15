package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"heytobi.dev/record-store-service/pkg/recordstore"
)

type Endpoints struct {
	CreateRecordsEndpoint endpoint.Endpoint
	GetRecordsEndpoint    endpoint.Endpoint
	SellRecordsEndpoint   endpoint.Endpoint
	DeleteRecordsEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc recordstore.Service) Endpoints {
	return Endpoints{
		CreateRecordsEndpoint: MakeCreateRecordsEndpoint(svc),
		GetRecordsEndpoint:    MakeGetRecordsEndpoint(svc),
		SellRecordsEndpoint:   MakeSellRecordsEndpoint(svc),
		DeleteRecordsEndpoint: MakeDeleteRecordsEndpoint(svc),
	}
}

func MakeCreateRecordsEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(CreateRecordsRequest)
		response, err := svc.CreateRecords(ctx)
		return CreateRecordsResponse{Response: response}, err
	}
}

func MakeGetRecordsEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetRecordsRequest)
		response, err := svc.GetRecords(ctx)
		return GetRecordsResponse{Response: response}, err
	}
}

func MakeSellRecordsEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(SellRecordsRequest)
		response, err := svc.SellRecords(ctx)
		return SellRecordsResponse{Response: response}, err
	}
}

func MakeDeleteRecordsEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(DeleteRecordsRequest)
		response, err := svc.DeleteRecords(ctx)
		return DeleteRecordsResponse{Response: response}, err
	}
}
