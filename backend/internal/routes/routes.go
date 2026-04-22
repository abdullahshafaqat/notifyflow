package routes

import (
	"net/http"

	"github.com/abdullahshafaqat/notifyflow/internal/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/send", handlers.SendNotificationHandler)
}
