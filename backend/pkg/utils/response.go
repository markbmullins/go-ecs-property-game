// pkg/utils/response.go

package utils

import (
	"encoding/json"
	"net/http"
)

// Response is a standard structure for all HTTP responses.
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Optional field for extra data
}

// SendResponse is a helper function to send JSON responses.
func SendResponse(w http.ResponseWriter, status string, message string, data interface{}, httpStatusCode int) {
	response := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	json.NewEncoder(w).Encode(response)
}
