package transport

import (
	"context"
	"encoding/json"
	"net/http"

	gkhttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"heytobi.dev/record-store-service/internal/errors"
)

func NewHTTPHandler(e Endpoints) http.Handler {
	router := mux.NewRouter()
	v1SubRouter := router.PathPrefix("/api/v1").Subrouter()
	options := []gkhttp.ServerOption{
		gkhttp.ServerErrorEncoder(encodeError),
	}
	createArtistsRoutes(e, v1SubRouter, options)
	createRecordsRoutes(e, v1SubRouter, options)
	return router
}

func createArtistsRoutes(e Endpoints, r *mux.Router, so []gkhttp.ServerOption) {
	endpoint := "/artists"

	r.Methods("POST").Path(endpoint).Handler(gkhttp.NewServer(
		e.CreateArtistEndpoint,
		decodeCreateNewArtistRequest,
		encodeResponse,
		so...))

	r.Methods("GET").Path(endpoint).Handler(gkhttp.NewServer(
		e.GetArtistEndpoint,
		decodeGetArtistRequest,
		encodeResponse,
		so...))

	r.Methods("DELETE").Path(endpoint).Handler(gkhttp.NewServer(
		e.DeleteArtistEndpoint,
		decodeDeleteArtistRequest,
		encodeResponse,
		so...))
}

func createRecordsRoutes(e Endpoints, r *mux.Router, so []gkhttp.ServerOption) {
	endpoint := "/records"

	r.Methods("POST").Path(endpoint).Handler(gkhttp.NewServer(
		e.CreateRecordEndpoint,
		decodeCreateNewRecordRequest,
		encodeResponse,
		so...))

	r.Methods("GET").Path(endpoint).Handler(gkhttp.NewServer(
		e.GetRecordEndpoint,
		decodeGetRecordRequest,
		encodeResponse,
		so...))

	r.Methods("PUT").Path(endpoint).Handler(gkhttp.NewServer(
		e.SellRecordEndpoint,
		decodeSellRecordRequest,
		encodeResponse,
		so...))

	r.Methods("DELETE").Path(endpoint).Handler(gkhttp.NewServer(
		e.DeleteRecordEndpoint,
		decodeDeleteRecordRequest,
		encodeResponse,
		so...))
}

func decodeCreateNewArtistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	name := r.FormValue("name")
	if name == "" {
		return nil, &errors.BadRequest{Message: "Missing required parameter name"}
	}
	return CreateArtistRequest{Name: name}, nil
}

func decodeGetArtistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetArtistRequest{}, nil
}

func decodeDeleteArtistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return DeleteArtistRequest{}, nil
}

func decodeCreateNewRecordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return CreateRecordRequest{}, nil
}

func decodeGetRecordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetRecordRequest{}, nil
}

func decodeSellRecordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return SellRecordRequest{}, nil
}

func decodeDeleteRecordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return DeleteRecordRequest{}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		encodeError(ctx, err, w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, isBadRequest := err.(*errors.BadRequest); isBadRequest {
		w.WriteHeader(http.StatusBadRequest)
	}

	if _, isPreconditionFailed := err.(*errors.PreconditionFailed); isPreconditionFailed {
		w.WriteHeader(http.StatusPreconditionFailed)
	}

	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
