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
	repo     db.DB
	grpc     *grpcclient.Client
	interval time.Duration
}

func NewScheduler(repo db.DB, grpc *grpcclient.Client, interval time.Duration) *Scheduler {
	if interval <= 0 {
		interval = 5 * time.Second
	}

	return &Scheduler{
		repo:     repo,
		grpc:     grpc,
		interval: interval,
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

	log.Printf("scheduler: found %d due notifications", len(dueNotifications))

	for _, notification := range dueNotifications {
		log.Printf("scheduler: dispatching notification %s", notification.ID)
		if err := s.repo.UpdateStatus(ctx, notification.ID, "processing", notification.Retry); err != nil {
			log.Printf("scheduler: failed to mark notification %s as processing: %v", notification.ID, err)
			continue
		}

		status, err := s.grpc.Send(ctx, notification.ID, notification.To, notification.Message)
		if err != nil {
			log.Printf("scheduler: failed to dispatch notification %s: %v", notification.ID, err)
			_ = s.repo.UpdateStatus(ctx, notification.ID, "failed", notification.Retry)
			continue
		}

		if status != pb.Status_SUCCESS {
			log.Printf("scheduler: notification %s returned non-success status: %s", notification.ID, status.String())
			_ = s.repo.UpdateStatus(ctx, notification.ID, "failed", notification.Retry)
		}
	}
}
