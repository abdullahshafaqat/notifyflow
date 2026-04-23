package service

import (
	"context"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

func (s *serviceImpl) GetAll(ctx context.Context) ([]models.Notification, error) {
	return s.database.GetAll(ctx)
}

func (s *serviceImpl) GetByID(ctx context.Context, id string) (models.Notification, error) {
	return s.database.GetByID(ctx, id)
}

func (s *serviceImpl) GetFailed(ctx context.Context) ([]models.Notification, error) {
	return s.database.GetFailed(ctx)
}
