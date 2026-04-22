package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
)

func StartWorker(workerCount int, maxRetries int, processingDelay time.Duration, retryBackoff time.Duration) {
	rand.Seed(time.Now().UnixNano())
	if workerCount <= 0 {
		workerCount = 1
	}
	if maxRetries <= 0 {
		maxRetries = 1
	}
	if processingDelay < 0 {
		processingDelay = 0
	}
	if retryBackoff < 0 {
		retryBackoff = 0
	}

	for i := 0; i < workerCount; i++ {
		go func(id int) {
			for job := range NotificationQueue {
				retries := maxRetries

				for attempt := 1; attempt <= retries; attempt++ {
					updateNotificationStatus(job.ID, "pending", attempt)

					err := processNotification(job, id, processingDelay)

					if err == nil {
						updateNotificationStatus(job.ID, "success", attempt)
						break
					}

					if attempt == retries {
						updateNotificationStatus(job.ID, "failed", attempt)
						fmt.Printf("[Worker-%d] ❌ Final failure Job ID: %s\n", id, job.ID)
						break
					}

					fmt.Printf("[Worker-%d] Retry %d for Job ID: %s\n", id, attempt, job.ID)
					time.Sleep(retryBackoff)
				}
			}
		}(i)
	}
}

func processNotification(n models.Notification, workerID int, processingDelay time.Duration) error {
	time.Sleep(processingDelay)

	if rand.Intn(2) == 0 {
		fmt.Printf("[Worker-%d] ❌ Failed Job ID: %s\n", workerID, n.ID)
		return fmt.Errorf("failed to send")
	}

	return nil
}
