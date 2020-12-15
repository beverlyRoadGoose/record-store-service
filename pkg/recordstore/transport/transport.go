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
	recordsEndpoint := "/records"

	v1SubRouter.Methods("POST").Path(recordsEndpoint).Handler(gkHttp.NewServer(
		e.CreateRecordsEndpoint,
		decodeCreateNewRecordRequest,
		encodeResponse))

	v1SubRouter.Methods("GET").Path(recordsEndpoint).Handler(gkHttp.NewServer(
		e.GetRecordsEndpoint,
		decodeGetRecordsRequest,
		encodeResponse))

	v1SubRouter.Methods("PUT").Path(recordsEndpoint).Handler(gkHttp.NewServer(
		e.SellRecordsEndpoint,
		decodeSellRecordsRequest,
		encodeResponse))

	v1SubRouter.Methods("DELETE").Path(recordsEndpoint).Handler(gkHttp.NewServer(
		e.DeleteRecordsEndpoint,
		decodeDeleteRecordsRequest,
		encodeResponse))

	return router
}

func decodeCreateNewRecordRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return CreateRecordsRequest{}, nil
}

func decodeGetRecordsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetRecordsRequest{}, nil
}

func decodeSellRecordsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return SellRecordsRequest{}, nil
}

func decodeDeleteRecordsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return DeleteRecordsRequest{}, nil
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
