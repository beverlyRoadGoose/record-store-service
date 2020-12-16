package transport

import (
	"github.com/google/uuid"
	"heytobi.dev/record-store-service/internal/entities"
)

type CreateArtistRequest struct {
	Name string `json:"name"`
}

type CreateArtistResponse struct {
	Artist entities.Artist `json:"artist"`
	Error  string          `json:"error,omitempty"`
}

type GetArtistRequest struct {
}

type GetArtistResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type DeleteArtistRequest struct {
}

type DeleteArtistResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type CreateRecordRequest struct {
	ArtistId uuid.UUID `json:"artistId"`
	Name     string    `json:"name"`
	Year     uint      `json:"year"`
}

type CreateRecordResponse struct {
	Record entities.Record `json:"record"`
	Error  string          `json:"error,omitempty"`
}

type GetRecordRequest struct {
}

type GetRecordResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type SellRecordRequest struct {
}

type SellRecordResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type DeleteRecordRequest struct {
}

type DeleteRecordResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}
