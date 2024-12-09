// pkg/utils/response.go

package utils

import (
	"encoding/json"
	"net/http"
)

// Response is a standard structure for all HTTP responses.
type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"` // Optional field for extra data
}

// SendResponse is a helper function to send JSON responses.
func SendResponse(w http.ResponseWriter, statusCode int, message string, data interface{}, httpStatusCode int) {
	response := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}
