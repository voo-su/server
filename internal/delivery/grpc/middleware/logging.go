package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

func UnaryLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)

	grpclog.Infof("Request - Method: %s \t Duration: %s \t Error: %v \n", info.FullMethod, time.Since(start), err)

	return h, err
}

func StreamLoggingInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	err := handler(srv, ss)

	grpclog.Infof("Stream - Method: %s \t Duration: %s \t Error: %v \n", info.FullMethod, time.Since(start), err)

	return err
}
