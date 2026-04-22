package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"github.com/abdullahshafaqat/notifyflow/internal/services"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, statusCode int, payload Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

func SendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, Response{
			Status:  "error",
			Message: "only POST allowed",
		})
		return
	}

	var notif models.Notification
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&notif)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Status:  "error",
			Message: "invalid request body",
		})
		return
	}

	err = services.SendNotification(notif)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, Response{
			Status:  "error",
			Message: err.Error(),
		})
		return
	}

	writeJSON(w, http.StatusOK, Response{
		Status:  "success",
		Message: "Notification sent successfully",
	})
}
