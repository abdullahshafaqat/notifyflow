package db

import (
	"context"
	"log"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "notifyflow"
	collectionName = "notifications"
)

func notificationsCollection() *mongo.Collection {
	return Client.Database(databaseName).Collection(collectionName)
}

func InitCollections() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := notificationsCollection()

	// Create index on id field
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "id", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatal("Failed to create index:", err)
	}

	log.Println("Notifications collection initialized with indexes ✅")
}

func SaveNotification(n models.Notification) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := notificationsCollection().InsertOne(ctx, n)
	return err
}

func UpdateNotificationStatus(id, status string, retryCount int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": id}
	update := bson.M{
		"$set": bson.M{
			"status":      status,
			"retry_count": retryCount,
		},
	}

	_, err := notificationsCollection().UpdateOne(ctx, filter, update)
	return err
}
