package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"github.com/google/uuid"
)

func (s *serviceImpl) Send(ctx context.Context, n models.Notification) error {
	if n.To == "" || n.Message == "" {
		return errors.New("missing required fields")
	}

	n.ID = uuid.NewString()
	n.Status = "pending"
	n.Retry = 0

	if err := s.database.Save(ctx, n); err != nil {
		return fmt.Errorf("failed to save notification: %w", err)
	}

	select {
	case s.queue <- n:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
