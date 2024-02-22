package models

type StatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DataResponse struct {
	Data  any `json:"data"`
	Count int `json:"count"`
}
