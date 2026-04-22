package services

import "github.com/abdullahshafaqat/notifyflow/internal/models"

var NotificationQueue chan models.Notification

func InitQueue(bufferSize int) {
	if bufferSize <= 0 {
		bufferSize = 1000
	}

	NotificationQueue = make(chan models.Notification, bufferSize)
}
