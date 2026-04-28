package db

import (
	"context"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *dbImpl) GetScheduledDue(ctx context.Context, now time.Time) ([]models.Notification, error) {
	cursor, err := s.collection.Find(ctx, bson.M{
		"status":  "scheduled",
		"send_at": bson.M{"$lte": now},
	})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []models.Notification
	err = cursor.All(ctx, &results)
	return results, err
}
