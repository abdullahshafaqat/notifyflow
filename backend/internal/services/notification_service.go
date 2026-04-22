package services

import (
	"errors"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"github.com/google/uuid"
)

func SendNotification(n models.Notification) error {
	if n.To == "" || n.Message == "" {
		return errors.New("missing required fields")
	}

	n.ID = uuid.New().String()
	n.Status = "pending"
	n.RetryCount = 0

	if err := savePendingNotification(n); err != nil {
		return err
	}

	if NotificationQueue == nil {
		InitQueue(1000)
	}

	NotificationQueue <- n

	return nil
}
