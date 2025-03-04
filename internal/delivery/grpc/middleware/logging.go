package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

func LoggingServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)

	grpclog.Infof("Request - Method: %s \t Duration: %s \t Error: %v \n", info.FullMethod, time.Since(start), err)

	return h, err
}
