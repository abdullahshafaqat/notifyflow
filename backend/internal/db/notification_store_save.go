package db

import (
	"context"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

func (s *dbImpl) Save(ctx context.Context, n models.Notification) error {
	_, err := s.collection.InsertOne(ctx, n)
	return err
}
