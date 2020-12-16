package transport

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"heytobi.dev/record-store-service/internal/errors"
)

func validateCreateArtistRequest(r *http.Request) error {
	if name := r.FormValue("name"); name == "" {
		return &errors.BadRequest{Message: "missing required parameter name"}
	}
	return nil
}

func validateCreateRecordRequest(r *http.Request) error {
	artistId := r.FormValue("artistId")
	if artistId == "" {
		return &errors.BadRequest{Message: "missing required parameter artistId"}
	}
	if _, err := uuid.Parse(artistId); err != nil {
		return &errors.BadRequest{Message: "artist id is not a valid uuid"}
	}

	if name := r.FormValue("name"); name == "" {
		return &errors.BadRequest{Message: "missing required parameter name"}
	}

	year := r.FormValue("year")
	if year == "" {
		return &errors.BadRequest{Message: "missing required parameter year"}
	}
	if _, err := strconv.Atoi(year); err != nil {
		return &errors.BadRequest{Message: "invalid year"}
	}

	return nil
}
