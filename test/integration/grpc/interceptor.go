package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"voo.su/test"
)

func unaryInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	config := test.GetConfig()
	token := config.Grpc.Token

	md := metadata.Pairs("Authorization", "Bearer "+token)
	ctx = metadata.NewOutgoingContext(ctx, md)

	return invoker(ctx, method, req, reply, cc, opts...)
}
