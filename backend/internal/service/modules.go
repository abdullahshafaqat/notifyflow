package service

import (
	"context"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

type Service interface {
	Send(ctx context.Context, n models.Notification) error
	GetAll(ctx context.Context) ([]models.Notification, error)
	GetByID(ctx context.Context, id string) (models.Notification, error)
	GetFailed(ctx context.Context) ([]models.Notification, error)
	UpdateStatus(ctx context.Context, id, status string, retry int) error
	ProcessNotification(ctx context.Context, n models.Notification, processingDelay time.Duration) error
}

type serviceImpl struct {
	database db.DB
	queue    chan models.Notification
}

func NewService(database db.DB, queue chan models.Notification) Service {
	return &serviceImpl{database: database, queue: queue}
}

func InitService(database db.DB, queue chan models.Notification) Service {
	return NewService(database, queue)
}
