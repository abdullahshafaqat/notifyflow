package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *dbImpl) UpdateStatus(ctx context.Context, id, status string, retry int) error {
	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"status": status,
			"retry":  retry,
		},
	}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}
