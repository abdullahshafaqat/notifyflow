package service

import (
	"github.com/abdullahshafaqat/notifyflow/internal/models"
	pb "github.com/abdullahshafaqat/notifyflow/proto"
)

func ConvertToModel(req *pb.NotificationRequest) models.Notification {
	return models.Notification{
		ID:      req.Id,
		To:      req.To,
		Message: req.Message,
	}
}
