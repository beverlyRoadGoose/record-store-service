package transport

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"heytobi.dev/record-store-service/pkg/recordstore"
)

type Endpoints struct {
	CreateRecordEndpoint endpoint.Endpoint
	GetRecordEndpoint    endpoint.Endpoint
	SellRecordEndpoint   endpoint.Endpoint
	DeleteRecordEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc recordstore.Service) Endpoints {
	return Endpoints{
		CreateRecordEndpoint: MakeCreateRecordEndpoint(svc),
		GetRecordEndpoint:    MakeGetRecordEndpoint(svc),
		SellRecordEndpoint:   MakeSellRecordEndpoint(svc),
		DeleteRecordEndpoint: MakeDeleteRecordEndpoint(svc),
	}
}

func MakeCreateRecordEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(CreateRecordRequest)
		response, err := svc.CreateRecord(ctx)
		return CreateRecordResponse{Response: response}, err
	}
}

func MakeGetRecordEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetRecordRequest)
		response, err := svc.GetRecord(ctx)
		return GetRecordResponse{Response: response}, err
	}
}

func MakeSellRecordEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(SellRecordRequest)
		response, err := svc.SellRecord(ctx)
		return SellRecordResponse{Response: response}, err
	}
}

func MakeDeleteRecordEndpoint(svc recordstore.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(DeleteRecordRequest)
		response, err := svc.DeleteRecord(ctx)
		return DeleteRecordResponse{Response: response}, err
	}
}
