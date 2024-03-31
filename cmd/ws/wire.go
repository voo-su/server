//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"voo.su/internal/config"
	"voo.su/internal/logic"
	"voo.su/internal/provider"
	"voo.su/internal/repository/cache"
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

	provider.NewPostgresqlClient,
	provider.NewRedisClient,
	provider.NewFilesystem,
	provider.NewEmailClient,
	provider.NewProviders,

	router.NewRouter,

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

	logic.NewMessageForwardLogic,

	service.NewDialogService,
	service.NewGroupMemberService,
	service.NewContactService,
)

func Initialize(conf *config.Config) *AppProvider {
	panic(wire.Build(
		ProviderSet,
		cache.ProviderSet,
		handler.ProviderSet,
		event.ProviderSet,
		consume.ProviderSet,
	))
}
