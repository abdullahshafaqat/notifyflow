package db

import (
	"context"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *dbImpl) GetByID(ctx context.Context, id string) (models.Notification, error) {
	var result models.Notification
	err := s.collection.FindOne(ctx, bson.M{"id": id}).Decode(&result)
	return result, err
}
