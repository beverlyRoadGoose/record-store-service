package transport

import (
	"context"
	"encoding/json"
	"net/http"

	gkHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(e Endpoints) http.Handler {
	router := mux.NewRouter()
	v1SubRouter := router.PathPrefix("/api/v1").Subrouter()
	createArtistsRoutes(e, v1SubRouter)
	createRecordsRoutes(e, v1SubRouter)
	return router
}

func createArtistsRoutes(e Endpoints, r *mux.Router) {
	endpoint := "/artists"

	r.Methods("POST").Path(endpoint).Handler(gkHttp.NewServer(
		e.CreateArtistEndpoint,
		decodeCreateNewArtistRequest,
		encodeResponse))

	r.Methods("GET").Path(endpoint).Handler(gkHttp.NewServer(
		e.GetArtistEndpoint,
		decodeGetArtistRequest,
		encodeResponse))

	r.Methods("DELETE").Path(endpoint).Handler(gkHttp.NewServer(
		e.DeleteArtistEndpoint,
		decodeDeleteArtistRequest,
		encodeResponse))
}

func createRecordsRoutes(e Endpoints, r *mux.Router) {
	endpoint := "/records"

	r.Methods("POST").Path(endpoint).Handler(gkHttp.NewServer(
		e.CreateRecordEndpoint,
		decodeCreateNewRecordRequest,
		encodeResponse))

	r.Methods("GET").Path(endpoint).Handler(gkHttp.NewServer(
		e.GetRecordEndpoint,
		decodeGetRecordRequest,
		encodeResponse))

	r.Methods("PUT").Path(endpoint).Handler(gkHttp.NewServer(
		e.SellRecordEndpoint,
		decodeSellRecordRequest,
		encodeResponse))

	r.Methods("DELETE").Path(endpoint).Handler(gkHttp.NewServer(
		e.DeleteRecordEndpoint,
		decodeDeleteRecordRequest,
		encodeResponse))
}

func decodeCreateNewArtistRequest(_ context.Context, r *http.Request) (interface{}, error) {
	name := r.FormValue("name")
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
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
