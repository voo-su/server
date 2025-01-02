// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package middleware

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/pkg/locale"
	"voo.su/pkg/middleware"
)

type AuthMiddleware struct {
	Conf   *config.Config
	Locale locale.ILocale
}

func NewAuthMiddleware(
	conf *config.Config,
	locale locale.ILocale,
) *AuthMiddleware {
	return &AuthMiddleware{
		Conf:   conf,
		Locale: locale,
	}
}

type GrpcMethod struct {
	Name string
}

type GrpcMethodService struct {
	PublicMethods []*GrpcMethod
}

func NewGrpMethodsService() *GrpcMethodService {
	return &GrpcMethodService{
		PublicMethods: []*GrpcMethod{
			{
				Name: "/auth.AuthService/Login",
			},
			{
				Name: "/auth.AuthService/Verify",
			},
		},
	}
}

func (g *GrpcMethodService) IsPublicMethod(method string) bool {
	isPublic := false
	for _, route := range g.PublicMethods {
		if route.Name == method {
			isPublic = true
		}
	}
	return isPublic
}

func AuthorizationServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	methodService := GetGlobalService(GrpcMethodsServiceKey).(*GrpcMethodService)

	if methodService.IsPublicMethod(info.FullMethod) {
		return handler(ctx, req)
	}

	authService := GetGlobalService(AuthMiddlewareKey).(*AuthMiddleware)

	claims, err := middleware.GrpcToken(ctx, authService.Locale, constant.GuardGrpcAuth, authService.Conf.App.Jwt.Secret)
	if err != nil {
		fmt.Println(err)
		return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	}

	//uid, err := strconv.Atoi(claims.ID)
	//if err != nil {
	//	fmt.Println(err)
	//
	//	return nil, status.Errorf(codes.Unauthenticated, "unauthorized")
	//}

	ctx = context.WithValue(ctx, "user", claims)

	//ctx = context.WithValue(ctx, middleware.JWTSessionConst, middleware.JSession{
	//	Uid:       uid,
	//	Token:     "token", // TODO
	//	ExpiresAt: claims.ExpiresAt.Unix(),
	//})

	return handler(ctx, req)
}
