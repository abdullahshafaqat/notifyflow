package models

type Response struct {
	ID      string `json:"id,omitempty"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
