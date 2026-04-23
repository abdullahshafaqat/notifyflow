package db

import (
	"context"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationStore interface {
	Save(ctx context.Context, n models.Notification) error
	UpdateStatus(ctx context.Context, id, status string, retry int) error
	GetAll(ctx context.Context) ([]models.Notification, error)
	GetByID(ctx context.Context, id string) (models.Notification, error)
	GetFailed(ctx context.Context) ([]models.Notification, error)
}

type DB interface {
	Save(ctx context.Context, n models.Notification) error
	UpdateStatus(ctx context.Context, id, status string, retry int) error
	GetAll(ctx context.Context) ([]models.Notification, error)
	GetByID(ctx context.Context, id string) (models.Notification, error)
	GetFailed(ctx context.Context) ([]models.Notification, error)
}

type dbImpl struct {
	collection *mongo.Collection
}

func NewDB(client *mongo.Client, databaseName string) DB {
	if databaseName == "" {
		databaseName = "notifyflow"
	}

	return &dbImpl{
		collection: client.Database(databaseName).Collection("notifications"),
	}
}

func InitDB(client *mongo.Client, databaseName string) DB {
	return NewDB(client, databaseName)
}
