package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	MongoURI          string
	ServerPort        string
	Environment       string
	WorkerCount       int
	QueueBuffer       int
	ProcessingDelayMS int
	RetryBackoffMS    int
	MaxRetries        int
	ResendAPIKey      string
	EmailFrom         string
}

var AppConfig *Config

func LoadConfig() {
	mongoURI := getEnv("MONGO_URI", "mongodb://localhost:27017")
	serverPort := getEnv("SERVER_PORT", "8080")
	environment := getEnv("ENVIRONMENT", "development")
	workerCount := getEnvInt("WORKER_COUNT", 50)
	queueBuffer := getEnvInt("QUEUE_BUFFER", 100000)
	processingDelayMS := getEnvInt("PROCESSING_DELAY_MS", 10)
	retryBackoffMS := getEnvInt("RETRY_BACKOFF_MS", 50)
	maxRetries := getEnvInt("MAX_RETRIES", 3)
	resendAPIKey := getEnv("RESEND_API_KEY", "")
	emailFrom := getEnv("EMAIL_FROM", "")

	AppConfig = &Config{
		MongoURI:          mongoURI,
		ServerPort:        serverPort,
		Environment:       environment,
		WorkerCount:       workerCount,
		QueueBuffer:       queueBuffer,
		ProcessingDelayMS: processingDelayMS,
		RetryBackoffMS:    retryBackoffMS,
		MaxRetries:        maxRetries,
		ResendAPIKey:      resendAPIKey,
		EmailFrom:         emailFrom,
	}

	log.Printf("Config loaded - Env: %s, Port: %s, Mongo: %s, Workers: %d, Queue: %d, DelayMS: %d, RetryBackoffMS: %d, MaxRetries: %d\n",
		AppConfig.Environment,
		AppConfig.ServerPort,
		AppConfig.MongoURI,
		AppConfig.WorkerCount,
		AppConfig.QueueBuffer,
		AppConfig.ProcessingDelayMS,
		AppConfig.RetryBackoffMS,
		AppConfig.MaxRetries,
	)
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return parsed
}
