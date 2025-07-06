package middlewares

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

func WriteError(w http.ResponseWriter, statusCode int, errorCode string, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{
		ErrorCode: errorCode,
		Message:   message,
	})
}
