package middleware

import (
	"context"
	"google.golang.org/grpc"
)

type GrpcMethod struct {
	Name string
}

type GrpcMethodService struct {
	publicMethods []*GrpcMethod
}

func NewGrpMethodsService() *GrpcMethodService {
	return &GrpcMethodService{
		publicMethods: []*GrpcMethod{
			{
				Name: "/auth.AuthService/Login",
			},
		},
	}
}

func (s *GrpcMethodService) IsPublicMethod(method string) bool {
	isPublic := false
	for _, route := range s.publicMethods {
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

	authService := GetGlobalService(AuthServiceKey).(*TokenMiddleware)
	_, err := authService.ValidateToken(ctx)
	if err != nil {
		return nil, err
	}

	return handler(ctx, req)
}
