package db

import (
	"context"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *dbImpl) GetAll(ctx context.Context) ([]models.Notification, error) {
	cursor, err := s.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Notification
	err = cursor.All(ctx, &results)
	if results == nil {
		results = make([]models.Notification, 0)
	}
	return results, err
}
