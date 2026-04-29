package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/api"
	"github.com/abdullahshafaqat/notifyflow/internal/config"
	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/grpcclient"
	"github.com/abdullahshafaqat/notifyflow/internal/scheduler"
	"github.com/abdullahshafaqat/notifyflow/internal/service"
	"github.com/joho/godotenv"
)

const databaseName = "notifyflow"

func main() {
	_ = godotenv.Load()
	config.LoadConfig()
	db.ConnectMongo()

	handler, jobScheduler := buildDependencies()
	go jobScheduler.Start(context.Background())

	mux := http.NewServeMux()
	router := api.NewRouter(handler)
	router.DefineRoutes(mux)

	addr := ":" + config.AppConfig.ServerPort
	log.Printf("Server running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, corsMiddleware(mux)))
}

func buildDependencies() (*api.NotificationHandler, *scheduler.Scheduler) {
	database := db.InitDB(db.Client, databaseName)
	grpc, err := grpcclient.NewClient()
	if err != nil {
		log.Fatalf("failed to connect to gRPC worker: %v", err)
	}

	notificationService := service.InitService(database, grpc)
	notificationHandler, _ := api.InitAPI(notificationService)
	jobScheduler := scheduler.NewScheduler(
		database,
		grpc,
		5*time.Second,
		config.AppConfig.MaxRetries,
		time.Duration(config.AppConfig.RetryBackoffMS)*time.Millisecond,
	)
	return notificationHandler, jobScheduler
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
