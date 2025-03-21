package model

import "time"

type AccessGrpcStreamLog struct {
	FullMethod string     `ch:"full_method"`
	CreatedAt  *time.Time `ch:"created_at"`
}

func (AccessGrpcStreamLog) TableName() string {
	return "access_grpc_stream_logs"
}
