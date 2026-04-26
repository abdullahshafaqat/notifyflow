package service

import (
	"context"
	"fmt"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

type NotificationService interface {
	Process(ctx context.Context, n models.Notification) error
}

type notificationServiceImpl struct {
	repo db.DB
}

func NewNotificationService(repo db.DB) NotificationService {
	return &notificationServiceImpl{repo: repo}
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

	if n.To == "fail@test.com" {
		return fmt.Errorf("simulated failure")
	}

	return nil
}
