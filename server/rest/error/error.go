package error

import (
	"encoding/json"
	"net/http"
)

// HTTPError defines structure for an http error
type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func Wrap(message string, statusCode int, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	err := HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
	json.NewEncoder(w).Encode(err)
}
