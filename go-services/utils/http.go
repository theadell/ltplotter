package utils

import (
	"encoding/json"
	"ltplotter/models"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, errors map[string]string) {
	errorResponse := models.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		Errors:     errors,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResponse)
}
