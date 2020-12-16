package transport

type CreateArtistRequest struct {
}

type CreateArtistResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
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
}

type CreateRecordResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
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
