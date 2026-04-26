package grpcclient

import (
	"context"
	"time"

	pb "github.com/abdullahshafaqat/notifyflow/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.NotificationServiceClient
}

func NewClient() (*Client, error) {
	conn, err := grpc.Dial(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		client: pb.NewNotificationServiceClient(conn),
	}, nil
}

func (c *Client) Close() error {
	if c == nil || c.conn == nil {
		return nil
	}
	return c.conn.Close()
}

func (c *Client) Send(ctx context.Context, id, to, message string) (pb.Status, error) {
	requestCtx := ctx
	cancel := func() {}
	if requestCtx == nil {
		requestCtx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	}
	defer cancel()

	resp, err := c.client.SendNotification(
		requestCtx,
		&pb.NotificationRequest{
			Id:      id,
			To:      to,
			Message: message,
		},
	)
	if err != nil {
		return pb.Status_UNKNOWN, err
	}

	return resp.Status, nil
}
