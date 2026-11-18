package response

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON response with the given status code and data
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			// Log error in production, for now just return
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// Error writes a JSON error response
func Error(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, map[string]string{"error": message})
}

// Success writes a JSON success response
func Success(w http.ResponseWriter, data interface{}) {
	JSON(w, http.StatusOK, data)
}
