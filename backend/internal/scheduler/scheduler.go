package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/grpcclient"
	pb "github.com/abdullahshafaqat/notifyflow/proto"
)

type Scheduler struct {
	repo         db.DB
	grpc         *grpcclient.Client
	interval     time.Duration
	maxRetries   int
	retryBackoff time.Duration
}

func NewScheduler(repo db.DB, grpc *grpcclient.Client, interval time.Duration, maxRetries int, retryBackoff time.Duration) *Scheduler {
	if interval <= 0 {
		interval = 5 * time.Second
	}
	if maxRetries <= 0 {
		maxRetries = 1
	}
	if retryBackoff < 0 {
		retryBackoff = 0
	}

	return &Scheduler{
		repo:         repo,
		grpc:         grpc,
		interval:     interval,
		maxRetries:   maxRetries,
		retryBackoff: retryBackoff,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(s.interval)
	defer ticker.Stop()

	log.Printf("scheduler: started with interval=%s", s.interval)

	for {
		s.dispatchDue(ctx)

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func (s *Scheduler) dispatchDue(ctx context.Context) {
	dueNotifications, err := s.repo.GetScheduledDue(ctx, time.Now().UTC())
	if err != nil {
		log.Printf("scheduler: failed to fetch due notifications: %v", err)
		return
	}

	// found due notifications

	for _, notification := range dueNotifications {
		for attempt := notification.Retry + 1; attempt <= s.maxRetries; attempt++ {
			if err := s.repo.UpdateStatus(ctx, notification.ID, "processing", attempt); err != nil {
				log.Printf("scheduler: failed to mark notification %s as processing: %v", notification.ID, err)
				break
			}

			status, err := s.grpc.Send(ctx, notification.ID, notification.To, notification.Message)
			if err == nil && status == pb.Status_SUCCESS {
				if err := s.repo.UpdateStatus(ctx, notification.ID, "sent", attempt); err != nil {
					log.Printf("scheduler: failed to mark notification %s as sent: %v", notification.ID, err)
				}
				break
			}

			if err != nil {
				log.Printf("scheduler: failed to dispatch notification %s on attempt %d: %v", notification.ID, attempt, err)
				_ = s.repo.SetLastError(ctx, notification.ID, err.Error())
			} else {
				log.Printf("scheduler: notification %s returned non-success status on attempt %d: %s", notification.ID, attempt, status.String())
				_ = s.repo.SetLastError(ctx, notification.ID, status.String())
			}

			if attempt == s.maxRetries {
				if err := s.repo.UpdateStatus(ctx, notification.ID, "failed", attempt); err != nil {
					log.Printf("scheduler: failed to mark notification %s as failed: %v", notification.ID, err)
				}
				break
			}

			if err := s.repo.UpdateStatus(ctx, notification.ID, "retrying", attempt); err != nil {
				log.Printf("scheduler: failed to mark notification %s as retrying: %v", notification.ID, err)
				break
			}

			select {
			case <-time.After(s.retryBackoff):
			case <-ctx.Done():
				return
			}
		}
	}
}
