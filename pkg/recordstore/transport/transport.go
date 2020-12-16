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
		e.CreateRecordEndpoint,
		decodeCreateNewRecordRequest,
		encodeResponse))

	v1SubRouter.Methods("GET").Path(recordsEndpoint).Handler(gkHttp.NewServer(
		e.GetRecordEndpoint,
		decodeGetRecordRequest,
		encodeResponse))

	v1SubRouter.Methods("PUT").Path(recordsEndpoint).Handler(gkHttp.NewServer(
		e.SellRecordEndpoint,
		decodeSellRecordRequest,
		encodeResponse))

	v1SubRouter.Methods("DELETE").Path(recordsEndpoint).Handler(gkHttp.NewServer(
		e.DeleteRecordEndpoint,
		decodeDeleteRecordRequest,
		encodeResponse))

	return router
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
