package api

import (
	"net/http"

	"github.com/abdullahshafaqat/notifyflow/internal/service"
)

type Router interface {
	DefineRoutes(mux *http.ServeMux)
}

type routerImpl struct {
	notificationHandler *NotificationHandler
}

func NewRouter(notificationHandler *NotificationHandler) Router {
	return &routerImpl{notificationHandler: notificationHandler}
}

func InitAPI(notificationService service.Service) (*NotificationHandler, Router) {
	handler := NewNotificationHandler(notificationService)
	router := NewRouter(handler)
	return handler, router
}
