package grpcutil

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"strings"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/locale"
)

func GrpcToken(ctx context.Context, locale locale.ILocale, guard string, secret string) (*jwtutil.AuthClaims, *string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil, errors.New(locale.Localize("failed_to_fetch_metadata"))
	}

	tokens := md.Get("Authorization")
	token := strings.TrimPrefix(tokens[len(tokens)-1], "Bearer ")

	userClaims, err := jwtutil.Verify(locale, guard, secret, token)
	if err != nil {
		return nil, nil, err
	}

	return userClaims, &token, nil
}

func ClientIp(ctx context.Context) string {
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

func UserAgent(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("failed to extract data from context")
		return ""
	}

	userAgent := md.Get("user-agent")
	if len(userAgent) == 0 {
		userAgent[0] = "unknown"
	}

	return userAgent[0]
}

func UserId(ctx context.Context) int {
	session, ok := ctx.Value(jwtutil.JWTSession).(*jwtutil.JSession)
	if !ok {
		fmt.Println("failed to retrieve user from context")
		return 0
	}

	return session.Uid
}

func UserToken(ctx context.Context) string {
	session, ok := ctx.Value(jwtutil.JWTSession).(*jwtutil.JSession)
	if !ok {
		fmt.Println("failed to retrieve user from context")
		return ""
	}

	return session.Token
}
