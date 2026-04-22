package services

import (
	"fmt"

	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

func savePendingNotification(n models.Notification) error {
	if err := db.SaveNotification(n); err != nil {
		return fmt.Errorf("failed to save notification: %w", err)
	}
	return nil
}

func updateNotificationStatus(id, status string, retryCount int) {
	if err := db.UpdateNotificationStatus(id, status, retryCount); err != nil {
		fmt.Printf("[DB] Failed to update status for Job ID %s: %v\n", id, err)
	}
}
