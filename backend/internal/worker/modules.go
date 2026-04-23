package worker

import (
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"github.com/abdullahshafaqat/notifyflow/internal/service"
)

func InitWorker(
	notificationService service.Service,
	queue <-chan models.Notification,
	workerCount int,
	maxRetries int,
	processingDelay time.Duration,
	retryBackoff time.Duration,
) *Manager {
	return NewManager(
		notificationService,
		queue,
		workerCount,
		maxRetries,
		processingDelay,
		retryBackoff,
	)
}
