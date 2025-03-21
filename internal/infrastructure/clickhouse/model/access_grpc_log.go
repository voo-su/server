package model

import "time"

type AccessGrpcLog struct {
	FullMethod string     `ch:"full_method"`
	CreatedAt  *time.Time `ch:"created_at"`
}

func (AccessGrpcLog) TableName() string {
	return "access_grpc_logs"
}
