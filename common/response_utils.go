package common

import (
	"encoding/json"
	"net/http"
)

// SendErrorResponse formats the error response in a consistent manner
func SendErrorResponse(w http.ResponseWriter, code string, message string, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": map[string]string{
			"code":    code,
			"message": message,
		},
	})
}
