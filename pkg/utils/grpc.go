package utils

import (
	"context"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"strings"
)

func GetGrpcClientIp(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ipList := md["x-forwarded-for"]; len(ipList) > 0 {
			return ipList[0]
		}
	}

	if p, ok := peer.FromContext(ctx); ok {
		addr := p.Addr.String()
		return strings.Split(addr, ":")[0]
	}

	return ""
}
