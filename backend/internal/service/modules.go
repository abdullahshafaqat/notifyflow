package service

import (
	"context"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/grpcclient"
	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

type Service interface {
	Send(ctx context.Context, n models.Notification) (string, string, error)
	GetAll(ctx context.Context) ([]models.Notification, error)
	GetByID(ctx context.Context, id string) (models.Notification, error)
	GetFailed(ctx context.Context) ([]models.Notification, error)
	UpdateStatus(ctx context.Context, id, status string, retry int) error
	SetLastError(ctx context.Context, id string, lastError string) error
	ProcessNotification(ctx context.Context, n models.Notification, processingDelay time.Duration) error
}

type serviceImpl struct {
	database db.DB
	grpc     *grpcclient.Client
}

func NewService(database db.DB, grpc *grpcclient.Client) Service {
	return &serviceImpl{database: database, grpc: grpc}
}

func InitService(database db.DB, grpc *grpcclient.Client) Service {
	return NewService(database, grpc)
}

func (s *serviceImpl) SetLastError(ctx context.Context, id string, lastError string) error {
	return s.database.SetLastError(ctx, id, lastError)
}
