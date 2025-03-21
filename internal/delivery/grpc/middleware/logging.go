package middleware

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"time"
	clickhouseModel "voo.su/internal/infrastructure/clickhouse/model"
	clickhouseRepo "voo.su/internal/infrastructure/clickhouse/repository"
)

type LoggingMiddleware struct {
	AccessGrpcLogRepository       *clickhouseRepo.AccessGrpcLogRepository
	AccessGrpcStreamLogRepository *clickhouseRepo.AccessGrpcStreamLogRepository
}

func NewLoggingMiddleware(
	accessGrpcLogRepository *clickhouseRepo.AccessGrpcLogRepository,
	accessGrpcStreamLogRepository *clickhouseRepo.AccessGrpcStreamLogRepository,
) *LoggingMiddleware {
	return &LoggingMiddleware{
		AccessGrpcLogRepository:       accessGrpcLogRepository,
		AccessGrpcStreamLogRepository: accessGrpcStreamLogRepository,
	}
}

func (l *LoggingMiddleware) UnaryLoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	h, err := handler(ctx, req)

	log.Printf("Request - Method: %s \t Duration: %s \t Error: %v \n", info.FullMethod, time.Since(start), err)

	if err := l.AccessGrpcLogRepository.Create(context.Background(), &clickhouseModel.AccessGrpcLog{
		FullMethod: info.FullMethod,
	}); err != nil {
		log.Printf("Failed to save access grpc log: %s", err)
	}

	return h, err
}

func (l *LoggingMiddleware) StreamLoggingInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	err := handler(srv, ss)

	log.Printf("Stream - Method: %s \t Duration: %s \t Error: %v \n", info.FullMethod, time.Since(start), err)

	if err := l.AccessGrpcStreamLogRepository.Create(context.Background(), &clickhouseModel.AccessGrpcStreamLog{
		FullMethod: info.FullMethod,
	}); err != nil {
		log.Printf("Failed to save access grpc stream log: %s", err)
	}

	return err
}
