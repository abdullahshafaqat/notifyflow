package main

import (
	"context"
	"log"
	"net"

	"github.com/abdullahshafaqat/notifyflow/internal/config"
	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/email"
	"github.com/abdullahshafaqat/notifyflow/internal/service"
	pb "github.com/abdullahshafaqat/notifyflow/proto"
	"github.com/joho/godotenv"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
	svc service.NotificationService
}

func (s *server) SendNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	// received job - verbose, removed

	n := service.ConvertToModel(req)

	err := s.svc.Process(ctx, n)

	if err != nil {
		log.Printf("Job %s failed: %v\n", req.Id, err)
		return &pb.NotificationResponse{
			Status: pb.Status_FAILED,
		}, nil
	}

	
	return &pb.NotificationResponse{
		Status: pb.Status_SUCCESS,
	}, nil
}

func main() {
	_ = godotenv.Load()

	config.LoadConfig()

	db.ConnectMongo()

	database := db.InitDB(db.Client, "notifyflow")
	sender, err := email.NewResendSender(
		config.AppConfig.ResendAPIKey,
		config.AppConfig.EmailFrom,
	)
	if err != nil {
		log.Fatalf("failed to initialize Resend sender: %v", err)
	}

	svc := service.NewNotificationService(database, sender)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterNotificationServiceServer(grpcServer, &server{
		svc: svc,
	})

	log.Println("gRPC Worker running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
