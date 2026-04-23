package service

import (
	"context"
	"fmt"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

func (s *serviceImpl) ProcessNotification(ctx context.Context, n models.Notification, processingDelay time.Duration) error {
	if processingDelay > 0 {
		select {
		case <-time.After(processingDelay):
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	if n.To == "fail@test.com" {
		return fmt.Errorf("simulated failure")
	}

	return nil
}
