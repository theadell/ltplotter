package models

type ErrorResponse struct {
	StatusCode int               `json:"status_code"`
	Message    string            `json:"message"`
	Errors     map[string]string `json:"errors,omitempty"`
}
