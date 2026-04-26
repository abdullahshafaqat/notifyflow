package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	pb "github.com/abdullahshafaqat/notifyflow/proto"
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

	if s.grpc == nil {
		return errors.New("grpc client is not initialized")
	}

	status, err := s.grpc.Send(ctx, n.ID, n.To, n.Message)
	if err != nil {
		return fmt.Errorf("failed to send notification to worker: %w", err)
	}

	if status != pb.Status_SUCCESS {
		return fmt.Errorf("worker returned non-success status: %s", status.String())
	}

	return nil
}
