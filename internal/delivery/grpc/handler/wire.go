package handler

import (
	"github.com/google/wire"
	"voo.su/api/grpc/pb"
	"voo.su/internal/delivery/grpc/handler/chat"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(pb.UnimplementedAuthServiceServer), "*"),
	wire.Bind(new(pb.AuthServiceServer), new(*Auth)),

	wire.Struct(new(pb.UnimplementedAccountServiceServer), "*"),
	wire.Bind(new(pb.AccountServiceServer), new(*Account)),

	wire.Struct(new(pb.UnimplementedChatServiceServer), "*"),
	wire.Bind(new(pb.ChatServiceServer), new(*chat.Chat)),

	wire.Struct(new(pb.UnimplementedGroupChatServiceServer), "*"),
	wire.Bind(new(pb.GroupChatServiceServer), new(*GroupChat)),

	wire.Struct(new(pb.UnimplementedContactServiceServer), "*"),
	wire.Bind(new(pb.ContactServiceServer), new(*Contact)),

	wire.Struct(new(pb.UnimplementedUploadServiceServer), "*"),
	wire.Bind(new(pb.UploadServiceServer), new(*Upload)),

	NewAuthHandler,
	chat.NewChatHandler,
	NewGroupChatHandler,
	NewContactHandler,
	NewAccountHandler,
	NewUploadHandler,
)
