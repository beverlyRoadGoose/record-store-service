package transport

type CreateRecordsRequest struct {
}

type CreateRecordsResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type GetRecordsRequest struct {
}

type GetRecordsResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type SellRecordsRequest struct {
}

type SellRecordsResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}

type DeleteRecordsRequest struct {
}

type DeleteRecordsResponse struct {
	Response string `json:"response"`
	Error    string `json:"error,omitempty"`
}
