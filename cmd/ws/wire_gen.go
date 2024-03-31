// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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
	chat2 "voo.su/internal/transport/ws/consume/chat"
	"voo.su/internal/transport/ws/event"
	"voo.su/internal/transport/ws/event/chat"
	"voo.su/internal/transport/ws/handler"
	"voo.su/internal/transport/ws/process"
	"voo.su/internal/transport/ws/router"
)

// Injectors from wire.go:

func Initialize(conf *config.Config) *AppProvider {
	client := provider.NewRedisClient(conf)
	serverStorage := cache.NewSidStorage(client)
	clientStorage := cache.NewClientStorage(client, conf, serverStorage)
	roomStorage := cache.NewRoomStorage(client)
	db := provider.NewPostgresqlClient(conf)
	relation := cache.NewRelation(client)
	groupChatMember := repo.NewGroupMember(db, relation)
	source := repo.NewSource(db, client)
	groupChatMemberService := service.NewGroupMemberService(source, groupChatMember)
	sequence := cache.NewSequence(client)
	repoSequence := repo.NewSequence(db, sequence)
	messageForwardLogic := logic.NewMessageForwardLogic(db, repoSequence)
	split := repo.NewFileSplit(db)
	vote := cache.NewVote(client)
	messageVote := repo.NewMessageVote(db, vote)
	filesystem := provider.NewFilesystem(conf)
	unreadStorage := cache.NewUnreadStorage(client)
	messageStorage := cache.NewMessageStorage(client)
	message := repo.NewMessage(db)
	bot := repo.NewBot(db)
	messageService := &service.MessageService{
		Source:              source,
		MessageForwardLogic: messageForwardLogic,
		GroupChatMemberRepo: groupChatMember,
		SplitRepo:           split,
		MessageVoteRepo:     messageVote,
		Filesystem:          filesystem,
		UnreadStorage:       unreadStorage,
		MessageStorage:      messageStorage,
		ServerStorage:       serverStorage,
		ClientStorage:       clientStorage,
		Sequence:            repoSequence,
		DialogVoteCache:     vote,
		MessageRepo:         message,
		BotRepo:             bot,
	}
	chatHandler := chat.NewHandler(client, groupChatMemberService, messageService)
	chatEvent := &event.ChatEvent{
		Redis:                  client,
		Config:                 conf,
		RoomStorage:            roomStorage,
		GroupChatMemberRepo:    groupChatMember,
		GroupChatMemberService: groupChatMemberService,
		Handler:                chatHandler,
	}
	chatChannel := &handler.ChatChannel{
		Storage: clientStorage,
		Event:   chatEvent,
	}
	handlerHandler := &handler.Handler{
		Chat:   chatChannel,
		Config: conf,
	}
	jwtTokenStorage := cache.NewTokenSessionStorage(client)
	engine := router.NewRouter(conf, handlerHandler, jwtTokenStorage)
	healthSubscribe := process.NewHealthSubscribe(conf, serverStorage)
	dialog := repo.NewDialog(db)
	dialogService := service.NewDialogService(source, dialog, groupChatMember)
	contactRemark := cache.NewContactRemark(client)
	contact := repo.NewContact(db, contactRemark, relation)
	contactService := service.NewContactService(source, contact)
	handler2 := chat2.NewHandler(conf, clientStorage, roomStorage, dialogService, messageService, contactService, source)
	chatSubscribe := consume.NewChatSubscribe(handler2)
	messageSubscribe := process.NewMessageSubscribe(conf, client, chatSubscribe)
	subServers := &process.SubServers{
		HealthSubscribe:  healthSubscribe,
		MessageSubscribe: messageSubscribe,
	}
	server := process.NewServer(subServers)
	emailClient := provider.NewEmailClient(conf)
	providers := provider.NewProviders(emailClient)
	appProvider := &AppProvider{
		Config:    conf,
		Engine:    engine,
		Coroutine: server,
		Handler:   handlerHandler,
		Providers: providers,
	}
	return appProvider
}

// wire.go:

var ProviderSet = wire.NewSet(wire.Struct(new(service.MessageService), "*"), wire.Bind(new(service.MessageSendService), new(*service.MessageService)), wire.Struct(new(handler.Handler), "*"), wire.Struct(new(AppProvider), "*"), wire.Struct(new(process.SubServers), "*"), provider.NewPostgresqlClient, provider.NewRedisClient, provider.NewFilesystem, provider.NewEmailClient, provider.NewProviders, router.NewRouter, process.NewServer, process.NewHealthSubscribe, process.NewMessageSubscribe, repo.NewSource, repo.NewDialog, repo.NewMessage, repo.NewMessageVote, repo.NewGroupMember, repo.NewContact, repo.NewFileSplit, repo.NewSequence, repo.NewBot, logic.NewMessageForwardLogic, service.NewDialogService, service.NewGroupMemberService, service.NewContactService)
