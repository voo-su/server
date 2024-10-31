package grpc

import (
	"github.com/google/wire"
	authPb "voo.su/api/grpc/pb"
	"voo.su/internal/repository"
	"voo.su/internal/service"
	"voo.su/internal/transport/grpc/handler"
	"voo.su/internal/transport/grpc/middleware"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),

	wire.Struct(new(authPb.UnimplementedAuthServiceServer), "*"),
	wire.Bind(new(authPb.AuthServiceServer), new(*handler.AuthHandler)),

	handler.NewAuthHandler,

	middleware.NewTokenMiddleware,
	middleware.NewGrpMethodsService,

	service.ProviderSet,

	repository.ProviderSet,
)
