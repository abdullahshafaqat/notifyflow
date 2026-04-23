package api

import "net/http"

func (r *routerImpl) DefineRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/send", r.notificationHandler.SendNotification)
	mux.HandleFunc("/notifications", r.notificationHandler.GetAllNotifications)
	mux.HandleFunc("/notification", r.notificationHandler.GetNotificationByID)
	mux.HandleFunc("/failed", r.notificationHandler.GetFailedNotifications)
}
