package worker

import (
	"context"
	"log"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/models"
	"github.com/abdullahshafaqat/notifyflow/internal/service"
)

type Manager struct {
	service         service.Service
	queue           <-chan models.Notification
	workerCount     int
	maxRetries      int
	processingDelay time.Duration
	retryBackoff    time.Duration
}

func NewManager(
	svc service.Service,
	queue <-chan models.Notification,
	workerCount int,
	maxRetries int,
	processingDelay time.Duration,
	retryBackoff time.Duration,
) *Manager {
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

	return &Manager{
		service:         svc,
		queue:           queue,
		workerCount:     workerCount,
		maxRetries:      maxRetries,
		processingDelay: processingDelay,
		retryBackoff:    retryBackoff,
	}
}

func (m *Manager) Start(ctx context.Context) {
	for i := 0; i < m.workerCount; i++ {
		go m.runWorker(ctx, i)
	}
}

func (m *Manager) runWorker(ctx context.Context, workerID int) {
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-m.queue:
			if !ok {
				return
			}
			m.processJob(ctx, workerID, job)
		}
	}
}

func (m *Manager) processJob(ctx context.Context, workerID int, job models.Notification) {
	for attempt := 1; attempt <= m.maxRetries; attempt++ {
		err := m.service.ProcessNotification(ctx, job, m.processingDelay)
		if err == nil {
			_ = m.service.UpdateStatus(ctx, job.ID, "success", attempt)
			return
		}

		if attempt == m.maxRetries {
			_ = m.service.UpdateStatus(ctx, job.ID, "failed", m.maxRetries)
			log.Printf("worker=%d job=%s state=failed attempts=%d", workerID, job.ID, m.maxRetries)
			return
		}

		_ = m.service.UpdateStatus(ctx, job.ID, "retrying", attempt)

		select {
		case <-time.After(m.retryBackoff):
		case <-ctx.Done():
			return
		}
	}
}
