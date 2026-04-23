package api

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

func writeJSON(w http.ResponseWriter, statusCode int, payload models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, statusCode int, message string) {
	writeJSON(w, statusCode, models.Response{Status: "error", Message: message})
}
