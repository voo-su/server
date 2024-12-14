package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

func ConnectToGRPC(address string) (*grpc.ClientConn, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Не удалось подключиться к gRPC серверу на %s: %v", address, err)
	}

	return conn, ctx, cancel
}
