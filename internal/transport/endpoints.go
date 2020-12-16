package transport

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"heytobi.dev/record-store-service/internal"
)

type Endpoints struct {
	CreateArtistEndpoint endpoint.Endpoint
	GetArtistEndpoint    endpoint.Endpoint
	DeleteArtistEndpoint endpoint.Endpoint
	CreateRecordEndpoint endpoint.Endpoint
	GetRecordEndpoint    endpoint.Endpoint
	SellRecordEndpoint   endpoint.Endpoint
	DeleteRecordEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc internal.Service) Endpoints {
	return Endpoints{
		CreateArtistEndpoint: MakeCreateArtistEndpoint(svc),
		GetArtistEndpoint:    MakeGetArtistEndpoint(svc),
		DeleteArtistEndpoint: MakeDeleteArtistEndpoint(svc),
		CreateRecordEndpoint: MakeCreateRecordEndpoint(svc),
		GetRecordEndpoint:    MakeGetRecordEndpoint(svc),
		SellRecordEndpoint:   MakeSellRecordEndpoint(svc),
		DeleteRecordEndpoint: MakeDeleteRecordEndpoint(svc),
	}
}

func MakeCreateArtistEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(CreateArtistRequest)
		if !ok {
			return nil, errors.New("invalid request")
		}
		artist, _ := svc.CreateArtist(ctx, req.Name)
		return CreateArtistResponse{Artist: artist}, nil
	}
}

func MakeGetArtistEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetArtistRequest)
		response, err := svc.GetArtist(ctx)
		return GetArtistResponse{Response: response}, err
	}
}

func MakeDeleteArtistEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(DeleteArtistRequest)
		response, err := svc.DeleteArtist(ctx)
		return DeleteArtistResponse{Response: response}, err
	}
}

func MakeCreateRecordEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(CreateRecordRequest)
		response, err := svc.CreateRecord(ctx)
		return CreateRecordResponse{Response: response}, err
	}
}

func MakeGetRecordEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetRecordRequest)
		response, err := svc.GetRecord(ctx)
		return GetRecordResponse{Response: response}, err
	}
}

func MakeSellRecordEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(SellRecordRequest)
		response, err := svc.SellRecord(ctx)
		return SellRecordResponse{Response: response}, err
	}
}

func MakeDeleteRecordEndpoint(svc internal.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(DeleteRecordRequest)
		response, err := svc.DeleteRecord(ctx)
		return DeleteRecordResponse{Response: response}, err
	}
}
