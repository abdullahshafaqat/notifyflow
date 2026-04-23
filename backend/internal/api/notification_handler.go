package api

import (
	"github.com/abdullahshafaqat/notifyflow/internal/service"
)

type NotificationHandler struct {
	service service.Service
}

func NewNotificationHandler(s service.Service) *NotificationHandler {
	return &NotificationHandler{service: s}
}
