package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/api"
	"github.com/abdullahshafaqat/notifyflow/internal/config"
	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"github.com/abdullahshafaqat/notifyflow/internal/service"
	"github.com/abdullahshafaqat/notifyflow/internal/worker"
	"github.com/joho/godotenv"
)

const databaseName = "notifyflow"

func main() {
	_ = godotenv.Load()
	config.LoadConfig()
	db.ConnectMongo()

	handler, workerManager := buildDependencies()
	workerManager.Start(context.Background())

	mux := http.NewServeMux()
	router := api.NewRouter(handler)
	router.DefineRoutes(mux)

	addr := ":" + config.AppConfig.ServerPort
	log.Printf("Server running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func buildDependencies() (*api.NotificationHandler, *worker.Manager) {
	database := db.InitDB(db.Client, databaseName)
	queue := make(chan models.Notification, config.AppConfig.QueueBuffer)
	notificationService := service.InitService(database, queue)
	notificationHandler, _ := api.InitAPI(notificationService)

	workerManager := worker.InitWorker(
		notificationService,
		queue,
		config.AppConfig.WorkerCount,
		config.AppConfig.MaxRetries,
		time.Duration(config.AppConfig.ProcessingDelayMS)*time.Millisecond,
		time.Duration(config.AppConfig.RetryBackoffMS)*time.Millisecond,
	)

	return notificationHandler, workerManager
}
