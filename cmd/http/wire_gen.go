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
	"voo.su/internal/transport/http/handler"
	"voo.su/internal/transport/http/handler/v1"
	"voo.su/internal/transport/http/router"
)

// Injectors from wire.go:

func Initialize(conf *config.Config) *AppProvider {
	client := provider.NewRedisClient(conf)
	smsStorage := cache.NewSmsStorage(client)
	db := provider.NewPostgresqlClient(conf)
	contactRemark := cache.NewContactRemark(client)
	relation := cache.NewRelation(client)
	contact := repo.NewContact(db, contactRemark, relation)
	groupChatMember := repo.NewGroupMember(db, relation)
	groupChat := repo.NewGroupChat(db)
	emailClient := provider.NewEmailClient(conf)
	user := repo.NewUser(db)
	authService := service.NewAuthService(smsStorage, contact, groupChatMember, groupChat, conf, emailClient, user)
	jwtTokenStorage := cache.NewTokenSessionStorage(client)
	redisLock := cache.NewRedisLock(client)
	source := repo.NewSource(db, client)
	httpClient := provider.NewHttpClient()
	requestClient := provider.NewRequestClient(httpClient)
	ipAddressService := service.NewIpAddressService(source, conf, requestClient)
	dialog := repo.NewDialog(db)
	dialogService := service.NewDialogService(source, dialog, groupChatMember)
	sequence := cache.NewSequence(client)
	repoSequence := repo.NewSequence(db, sequence)
	messageForwardLogic := logic.NewMessageForwardLogic(db, repoSequence)
	split := repo.NewFileSplit(db)
	vote := cache.NewVote(client)
	messageVote := repo.NewMessageVote(db, vote)
	filesystem := provider.NewFilesystem(conf)
	unreadStorage := cache.NewUnreadStorage(client)
	messageStorage := cache.NewMessageStorage(client)
	serverStorage := cache.NewSidStorage(client)
	clientStorage := cache.NewClientStorage(client, conf, serverStorage)
	message := repo.NewMessage(db)
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
	}
	userSession := repo.NewUserSession(db)
	auth := &v1.Auth{
		Config:             conf,
		AuthService:        authService,
		JwtTokenStorage:    jwtTokenStorage,
		RedisLock:          redisLock,
		IpAddressService:   ipAddressService,
		DialogService:      dialogService,
		MessageSendService: messageService,
		UserSession:        userSession,
	}
	account := &v1.Account{
		UserRepo: user,
	}
	v1User := &v1.User{
		UserRepo: user,
	}
	contactService := service.NewContactService(source, contact)
	v1Contact := &v1.Contact{
		ContactService:     contactService,
		ClientStorage:      clientStorage,
		DialogService:      dialogService,
		MessageSendService: messageService,
		ContactRepo:        contact,
		UserRepo:           user,
		DialogRepo:         dialog,
	}
	contactRequestService := service.NewContactApplyService(source)
	contactRequest := &v1.ContactRequest{
		ContactRequestService: contactRequestService,
		ContactService:        contactService,
		MessageSendService:    messageService,
		ContactRepo:           contact,
	}
	groupChatService := service.NewGroupChatService(source, groupChat, groupChatMember, relation, repoSequence)
	v1Dialog := &v1.Dialog{
		DialogService:    dialogService,
		RedisLock:        redisLock,
		ClientStorage:    clientStorage,
		MessageStorage:   messageStorage,
		ContactService:   contactService,
		UnreadStorage:    unreadStorage,
		ContactRemark:    contactRemark,
		GroupChatService: groupChatService,
		AuthService:      authService,
		ContactRepo:      contact,
		UserRepo:         user,
		GroupChatRepo:    groupChat,
	}
	groupChatMemberService := service.NewGroupMemberService(source, groupChatMember)
	v1Message := &v1.Message{
		DialogService:          dialogService,
		AuthService:            authService,
		MessageSendService:     messageService,
		Filesystem:             filesystem,
		MessageService:         messageService,
		GroupChatMemberService: groupChatMemberService,
		GroupMemberRepo:        groupChatMember,
		MessageRepo:            message,
	}
	publish := &v1.Publish{
		AuthService:        authService,
		MessageSendService: messageService,
	}
	splitService := service.NewSplitService(source, split, conf, filesystem)
	upload := &v1.Upload{
		Config:       conf,
		Filesystem:   filesystem,
		SplitService: splitService,
	}
	groupChatNotice := repo.NewGroupChatNotice(db)
	groupChatNoticeService := service.NewGroupChatNoticeService(source, groupChatNotice)
	v1GroupChat := &v1.GroupChat{
		Repo:                   source,
		UserRepo:               user,
		GroupChatRepo:          groupChat,
		GroupChatMemberRepo:    groupChatMember,
		DialogRepo:             dialog,
		GroupChatService:       groupChatService,
		GroupChatMemberService: groupChatMemberService,
		DialogService:          dialogService,
		ContactService:         contactService,
		GroupNoticeService:     groupChatNoticeService,
		MessageSendService:     messageService,
		RedisLock:              redisLock,
	}
	groupChatRequestStorage := cache.NewGroupChatRequestStorage(client)
	groupChatRequest := repo.NewGroupChatApply(db)
	groupChatRequestService := service.NewGroupRequestService(source, groupChatRequest)
	v1GroupChatRequest := &v1.GroupChatRequest{
		GroupRequestStorage:     groupChatRequestStorage,
		GroupChatRepo:           groupChat,
		GroupChatRequestRepo:    groupChatRequest,
		GroupMemberRepo:         groupChatMember,
		GroupChatRequestService: groupChatRequestService,
		GroupChatMemberService:  groupChatMemberService,
		GroupChatService:        groupChatService,
		Redis:                   client,
	}
	handlerV1 := &handler.V1{
		Auth:             auth,
		Account:          account,
		User:             v1User,
		Contact:          v1Contact,
		ContactRequest:   contactRequest,
		Dialog:           v1Dialog,
		Message:          v1Message,
		MessagePublish:   publish,
		Upload:           upload,
		GroupChat:        v1GroupChat,
		GroupChatRequest: v1GroupChatRequest,
	}
	handlerHandler := &handler.Handler{
		V1: handlerV1,
	}
	engine := router.NewRouter(conf, handlerHandler, jwtTokenStorage)
	appProvider := &AppProvider{
		Config: conf,
		Engine: engine,
	}
	return appProvider
}

// wire.go:

var ProviderSet = wire.NewSet(wire.Struct(new(AppProvider), "*"), router.NewRouter, provider.NewPostgresqlClient, provider.NewRedisClient, provider.NewHttpClient, provider.NewEmailClient, provider.NewFilesystem, provider.NewRequestClient, wire.Struct(new(handler.Handler), "*"))
