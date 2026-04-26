package main

import (
	"log"
	"net/http"

	"github.com/abdullahshafaqat/notifyflow/internal/api"
	"github.com/abdullahshafaqat/notifyflow/internal/config"
	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/grpcclient"
	"github.com/abdullahshafaqat/notifyflow/internal/service"
	"github.com/joho/godotenv"
)

const databaseName = "notifyflow"

func main() {
	_ = godotenv.Load()
	config.LoadConfig()
	db.ConnectMongo()

	handler := buildDependencies()

	mux := http.NewServeMux()
	router := api.NewRouter(handler)
	router.DefineRoutes(mux)

	addr := ":" + config.AppConfig.ServerPort
	log.Printf("Server running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func buildDependencies() *api.NotificationHandler {
	database := db.InitDB(db.Client, databaseName)
	grpc, err := grpcclient.NewClient()
	if err != nil {
		log.Fatalf("failed to connect to gRPC worker: %v", err)
	}

	notificationService := service.InitService(database, grpc)
	notificationHandler, _ := api.InitAPI(notificationService)
	return notificationHandler
}
