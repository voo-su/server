package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strconv"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/pkg/grpcutil"
	"voo.su/pkg/jwtutil"
	"voo.su/pkg/locale"
)

type AuthMiddleware struct {
	Conf          *config.Config
	Locale        locale.ILocale
	PublicMethods map[string]struct{}
}

var publicMethods = map[string]struct{}{
	"/auth.AuthService/Login":  {},
	"/auth.AuthService/Verify": {},
}

func NewAuthMiddleware(conf *config.Config, locale locale.ILocale) *AuthMiddleware {
	return &AuthMiddleware{
		Conf:          conf,
		Locale:        locale,
		PublicMethods: publicMethods,
	}
}

func (a *AuthMiddleware) IsPublicMethod(method string) bool {
	_, exists := a.PublicMethods[method]
	return exists
}

func (a *AuthMiddleware) validateToken(ctx context.Context) (*jwtutil.JSession, error) {
	claims, token, err := grpcutil.GrpcToken(ctx, a.Locale, constant.GuardGrpcAuth, a.Conf.App.Jwt.Secret)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	uid, err := strconv.Atoi(claims.ID)
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	return &jwtutil.JSession{
		Uid:       uid,
		Token:     *token,
		ExpiresAt: claims.ExpiresAt.Unix(),
	}, nil
}

func (a *AuthMiddleware) UnaryAuthInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	if a.IsPublicMethod(info.FullMethod) {
		return handler(ctx, req)
	}

	claims, err := a.validateToken(ctx)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, jwtutil.JWTSession, claims)

	return handler(ctx, req)
}

type wrappedStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

func (a *AuthMiddleware) StreamAuthInterceptor(srv any, ss grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	claims, err := a.validateToken(ss.Context())
	if err != nil {
		return err
	}

	ctx := context.WithValue(ss.Context(), jwtutil.JWTSession, claims)

	wrapped := &wrappedStream{ServerStream: ss, ctx: ctx}
	return handler(srv, wrapped)
}
