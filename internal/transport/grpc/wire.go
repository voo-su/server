package grpc

import (
	"github.com/google/wire"
	"voo.su/api/grpc/pb"
	"voo.su/internal/repository"
	"voo.su/internal/service"
	"voo.su/internal/transport/grpc/handler"
	"voo.su/internal/transport/grpc/middleware"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),

	wire.Struct(new(pb.UnimplementedAuthServiceServer), "*"),
	wire.Bind(new(pb.AuthServiceServer), new(*handler.AuthHandler)),

	wire.Struct(new(pb.UnimplementedChatServiceServer), "*"),
	wire.Bind(new(pb.ChatServiceServer), new(*handler.ChatHandler)),

	wire.Struct(new(pb.UnimplementedContactServiceServer), "*"),
	wire.Bind(new(pb.ContactServiceServer), new(*handler.ContactHandler)),

	handler.NewAuthHandler,
	handler.NewChatHandler,
	handler.NewContactHandler,

	middleware.NewTokenMiddleware,
	middleware.NewGrpMethodsService,

	service.ProviderSet,

	repository.ProviderSet,
)
