package ws

import (
	"github.com/google/wire"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/internal/transport/ws/consume"
	"voo.su/internal/transport/ws/event"
	"voo.su/internal/transport/ws/handler"
	"voo.su/internal/transport/ws/process"
	"voo.su/internal/transport/ws/router"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(service.MessageService), "*"),
	wire.Bind(new(service.MessageSendService), new(*service.MessageService)),
	wire.Struct(new(handler.Handler), "*"),
	wire.Struct(new(AppProvider), "*"),
	wire.Struct(new(process.SubServers), "*"),

	router.NewRouter,

	service.NewDialogService,
	service.NewGroupMemberService,
	service.NewContactService,

	handler.ProviderSet,
	event.ProviderSet,
	consume.ProviderSet,

	process.NewServer,
	process.NewHealthSubscribe,
	process.NewMessageSubscribe,

	repo.NewSource,
	repo.NewDialog,
	repo.NewMessage,
	repo.NewMessageVote,
	repo.NewGroupMember,
	repo.NewContact,
	repo.NewFileSplit,
	repo.NewSequence,
	repo.NewBot,
)
