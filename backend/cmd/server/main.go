package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/config"
	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/routes"
	"github.com/abdullahshafaqat/notifyflow/internal/services"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.LoadConfig()

	services.InitQueue(config.AppConfig.QueueBuffer)
	db.ConnectMongo()

	services.StartWorker(
		config.AppConfig.WorkerCount,
		config.AppConfig.MaxRetries,
		time.Duration(config.AppConfig.ProcessingDelayMS)*time.Millisecond,
		time.Duration(config.AppConfig.RetryBackoffMS)*time.Millisecond,
	)

	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	addr := ":" + config.AppConfig.ServerPort
	fmt.Println("Server running on http://localhost" + addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
