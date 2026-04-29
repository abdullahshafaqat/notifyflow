package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *dbImpl) SetLastError(ctx context.Context, id string, lastError string) error {
	filter := bson.M{"id": id}
	update := bson.M{"$set": bson.M{"last_error": lastError}}

	_, err := s.collection.UpdateOne(ctx, filter, update)
	return err
}
