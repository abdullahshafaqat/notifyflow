package api

import (
	"encoding/json"
	"net/http"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

func (h *NotificationHandler) SendNotification(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "only POST allowed")
		return
	}

	var notif models.Notification
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&notif); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	id, status, err := h.service.Send(r.Context(), notif)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	message := "Notification sent successfully"
	if status == "scheduled" {
		message = "Notification scheduled successfully"
	}

	writeJSON(w, http.StatusOK, models.Response{ID: id, Status: status, Message: message})
}
