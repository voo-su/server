package grpc

import (
	"github.com/google/wire"
	"voo.su/api/grpc/gen/go/pb"
	"voo.su/internal/delivery/grpc/handler"
	"voo.su/internal/delivery/grpc/handler/chat"
	"voo.su/internal/delivery/grpc/middleware"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),

	wire.Struct(new(pb.UnimplementedAuthServiceServer), "*"),
	wire.Bind(new(pb.AuthServiceServer), new(*handler.Auth)),

	wire.Struct(new(pb.UnimplementedAccountServiceServer), "*"),
	wire.Bind(new(pb.AccountServiceServer), new(*handler.Account)),

	wire.Struct(new(pb.UnimplementedChatServiceServer), "*"),
	wire.Bind(new(pb.ChatServiceServer), new(*chat.Chat)),

	wire.Struct(new(pb.UnimplementedGroupChatServiceServer), "*"),
	wire.Bind(new(pb.GroupChatServiceServer), new(*handler.GroupChat)),

	wire.Struct(new(pb.UnimplementedContactServiceServer), "*"),
	wire.Bind(new(pb.ContactServiceServer), new(*handler.Contact)),

	handler.NewAuthHandler,
	chat.NewChatHandler,
	handler.NewGroupChatHandler,
	handler.NewContactHandler,
	handler.NewAccountHandler,

	middleware.NewAuthMiddleware,
	middleware.NewGrpMethodsService,
)
