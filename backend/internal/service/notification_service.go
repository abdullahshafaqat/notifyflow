package service

import (
	"context"
	"fmt"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/email"
	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

type NotificationService interface {
	Process(ctx context.Context, n models.Notification) error
}

type notificationServiceImpl struct {
	repo   db.DB
	sender email.Sender
}

func NewNotificationService(repo db.DB, sender email.Sender) NotificationService {
	return &notificationServiceImpl{repo: repo, sender: sender}
}

func (s *notificationServiceImpl) Process(ctx context.Context, n models.Notification) error {

	maxRetries := 3
	retryBackoff := 25 * time.Millisecond
	processingDelay := 10 * time.Millisecond

	for attempt := 1; attempt <= maxRetries; attempt++ {
		err := s.processNotification(ctx, n, processingDelay)
		if err == nil {

			return s.repo.UpdateStatus(ctx, n.ID, "success", attempt)
		}

		if attempt == maxRetries {

			return s.repo.UpdateStatus(ctx, n.ID, "failed", maxRetries)
		}

		s.repo.UpdateStatus(ctx, n.ID, "retrying", attempt)
		time.Sleep(retryBackoff)
	}

	return nil
}

func (s *notificationServiceImpl) processNotification(ctx context.Context, n models.Notification, delay time.Duration) error {
	if delay > 0 {
		select {
		case <-time.After(delay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	if s.sender == nil {
		return fmt.Errorf("email sender is not initialized")
	}

	err := s.sender.Send(ctx, n.To, "NotifyFlow Notification", n.Message)
	if err != nil {
		return err
	}

	return nil
}
