package main

import (
	"context"
	"log"
	"net"

	"github.com/abdullahshafaqat/notifyflow/internal/config"
	"github.com/abdullahshafaqat/notifyflow/internal/db"
	"github.com/abdullahshafaqat/notifyflow/internal/service"
	pb "github.com/abdullahshafaqat/notifyflow/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedNotificationServiceServer
	svc service.NotificationService
}

func (s *server) SendNotification(ctx context.Context, req *pb.NotificationRequest) (*pb.NotificationResponse, error) {
	log.Println("Received job:", req.Id)

	// convert proto → model
	n := service.ConvertToModel(req)

	// process using existing logic
	err := s.svc.Process(ctx, n)

	if err != nil {
		log.Printf("Job %s failed: %v\n", req.Id, err)
		return &pb.NotificationResponse{
			Status: "failed",
		}, nil
	}

	log.Printf("Job %s completed successfully\n", req.Id)
	return &pb.NotificationResponse{
		Status: "success",
	}, nil
}

func main() {
	// Load config
	config.LoadConfig()

	// connect DB
	db.ConnectMongo()

	// init repo + service
	database := db.InitDB(db.Client, "notifyflow")
	svc := service.NewNotificationService(database)

	// start gRPC server
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
