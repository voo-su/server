package handler

import (
	"context"
	authPb "voo.su/api/grpc/pb"
	"voo.su/internal/transport/grpc/middleware"
)

type AuthHandler struct {
	authPb.UnimplementedAuthServiceServer
	TokenMiddleware *middleware.TokenMiddleware
}

func NewAuthHandler(
	tokenMiddleware *middleware.TokenMiddleware,
) *AuthHandler {
	return &AuthHandler{
		TokenMiddleware: tokenMiddleware,
	}
}
func (a *AuthHandler) Login(ctx context.Context, in *authPb.AuthLoginRequest) (*authPb.AuthLoginResponse, error) {

	return &authPb.AuthLoginResponse{}, nil
}
