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
	"voo.su/internal/transport/cli"
	"voo.su/internal/transport/cli/handle/cron"
	"voo.su/internal/transport/cli/handle/queue"
	"voo.su/internal/transport/http"
	"voo.su/internal/transport/http/handler"
	"voo.su/internal/transport/http/handler/v1"
	"voo.su/internal/transport/http/router"
	"voo.su/internal/transport/ws"
	"voo.su/internal/transport/ws/consume"
	chat2 "voo.su/internal/transport/ws/consume/chat"
	"voo.su/internal/transport/ws/event"
	"voo.su/internal/transport/ws/event/chat"
	handler2 "voo.su/internal/transport/ws/handler"
	"voo.su/internal/transport/ws/process"
	router2 "voo.su/internal/transport/ws/router"
)

// Injectors from wire.go:

func NewHttpInjector(conf *config.Config) *http.AppProvider {
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
	chatService := service.NewChatService(source, dialog, groupChatMember)
	bot := repo.NewBot(db)
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
	clientStorage := cache.NewClientStorage(conf, client, serverStorage)
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
		BotRepo:             bot,
	}
	userSession := repo.NewUserSession(db)
	auth := &v1.Auth{
		Conf:               conf,
		AuthService:        authService,
		JwtTokenStorage:    jwtTokenStorage,
		RedisLock:          redisLock,
		IpAddressService:   ipAddressService,
		ChatService:        chatService,
		BotRepo:            bot,
		MessageSendService: messageService,
		UserSession:        userSession,
	}
	account := &v1.Account{
		UserRepo: user,
	}
	contactService := service.NewContactService(source, contact)
	v1Contact := &v1.Contact{
		ContactService:     contactService,
		ClientStorage:      clientStorage,
		ChatService:        chatService,
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
	chat := &v1.Chat{
		ChatService:      chatService,
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
		ChatService:            chatService,
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
		Conf:         conf,
		Filesystem:   filesystem,
		SplitService: splitService,
	}
	v1GroupChat := &v1.GroupChat{
		Repo:                   source,
		UserRepo:               user,
		GroupChatRepo:          groupChat,
		GroupChatMemberRepo:    groupChatMember,
		DialogRepo:             dialog,
		GroupChatService:       groupChatService,
		GroupChatMemberService: groupChatMemberService,
		ChatService:            chatService,
		ContactService:         contactService,
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
	sticker := repo.NewSticker(db)
	stickerService := service.NewStickerService(source, sticker, filesystem)
	v1Sticker := &v1.Sticker{
		StickerRepo:    sticker,
		Filesystem:     filesystem,
		StickerService: stickerService,
		RedisLock:      redisLock,
	}
	contactFolder := repo.NewContactFolder(db)
	contactFolderService := service.NewContactFolderService(source, contact, contactFolder)
	v1ContactFolder := &v1.ContactFolder{
		ContactFolderService: contactFolderService,
	}
	groupChatAds := repo.NewGroupChatAds(db)
	groupChatAdsService := service.NewGroupChatAdsService(source, groupChatAds)
	v1GroupChatAds := &v1.GroupChatAds{
		GroupAdsService:     groupChatAdsService,
		GroupMemberService:  groupChatMemberService,
		MessageSendService:  messageService,
		GroupChatMemberRepo: groupChatMember,
		GroupChatAdsRepo:    groupChatAds,
	}
	search := &v1.Search{
		UserRepo:            user,
		GroupChatRepo:       groupChat,
		GroupChatMemberRepo: groupChatMember,
	}
	handlerV1 := &handler.V1{
		Auth:             auth,
		Account:          account,
		Contact:          v1Contact,
		ContactRequest:   contactRequest,
		Chat:             chat,
		Message:          v1Message,
		MessagePublish:   publish,
		Upload:           upload,
		GroupChat:        v1GroupChat,
		GroupChatRequest: v1GroupChatRequest,
		Sticker:          v1Sticker,
		ContactFolder:    v1ContactFolder,
		GroupChatAds:     v1GroupChatAds,
		Search:           search,
	}
	handlerHandler := &handler.Handler{
		V1: handlerV1,
	}
	engine := router.NewRouter(conf, handlerHandler, jwtTokenStorage)
	appProvider := &http.AppProvider{
		Conf:   conf,
		Engine: engine,
	}
	return appProvider
}

func NewWsInjector(conf *config.Config) *ws.AppProvider {
	client := provider.NewRedisClient(conf)
	serverStorage := cache.NewSidStorage(client)
	clientStorage := cache.NewClientStorage(conf, client, serverStorage)
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
		Conf:                   conf,
		RoomStorage:            roomStorage,
		GroupChatMemberRepo:    groupChatMember,
		GroupChatMemberService: groupChatMemberService,
		Handler:                chatHandler,
	}
	chatChannel := &handler2.ChatChannel{
		Storage: clientStorage,
		Event:   chatEvent,
	}
	handlerHandler := &handler2.Handler{
		Chat: chatChannel,
		Conf: conf,
	}
	jwtTokenStorage := cache.NewTokenSessionStorage(client)
	engine := router2.NewRouter(conf, handlerHandler, jwtTokenStorage)
	healthSubscribe := process.NewHealthSubscribe(conf, serverStorage)
	dialog := repo.NewDialog(db)
	chatService := service.NewChatService(source, dialog, groupChatMember)
	contactRemark := cache.NewContactRemark(client)
	contact := repo.NewContact(db, contactRemark, relation)
	contactService := service.NewContactService(source, contact)
	handler3 := chat2.NewHandler(conf, clientStorage, roomStorage, chatService, messageService, contactService, source)
	chatSubscribe := consume.NewChatSubscribe(handler3)
	messageSubscribe := process.NewMessageSubscribe(conf, client, chatSubscribe)
	subServers := &process.SubServers{
		HealthSubscribe:  healthSubscribe,
		MessageSubscribe: messageSubscribe,
	}
	server := process.NewServer(subServers)
	emailClient := provider.NewEmailClient(conf)
	providers := &provider.Providers{
		EmailClient: emailClient,
	}
	appProvider := &ws.AppProvider{
		Conf:      conf,
		Engine:    engine,
		Coroutine: server,
		Handler:   handlerHandler,
		Providers: providers,
	}
	return appProvider
}

func NewCronInjector(conf *config.Config) *cli.CronProvider {
	client := provider.NewRedisClient(conf)
	serverStorage := cache.NewSidStorage(client)
	clearWsCache := cron.NewClearWsCache(serverStorage)
	db := provider.NewPostgresqlClient(conf)
	filesystem := provider.NewFilesystem(conf)
	clearTmpFile := cron.NewClearTmpFile(db, filesystem)
	clearExpireServer := cron.NewClearExpireServer(serverStorage)
	crontab := &cli.Crontab{
		ClearWsCache:      clearWsCache,
		ClearTmpFile:      clearTmpFile,
		ClearExpireServer: clearExpireServer,
	}
	cronProvider := &cli.CronProvider{
		Conf:    conf,
		Crontab: crontab,
	}
	return cronProvider
}

func NewQueueInjector(conf *config.Config) *cli.QueueProvider {
	db := provider.NewPostgresqlClient(conf)
	client := provider.NewRedisClient(conf)
	emailHandle := queue.EmailHandle{
		Redis: client,
	}
	loginHandle := queue.LoginHandle{
		Redis: client,
	}
	queueJobs := &cli.QueueJobs{
		EmailHandle: emailHandle,
		LoginHandle: loginHandle,
	}
	queueProvider := &cli.QueueProvider{
		Conf: conf,
		DB:   db,
		Jobs: queueJobs,
	}
	return queueProvider
}

func NewMigrateInjector(conf *config.Config) *cli.MigrateProvider {
	db := provider.NewPostgresqlClient(conf)
	migrateProvider := &cli.MigrateProvider{
		Conf: conf,
		DB:   db,
	}
	return migrateProvider
}

// wire.go:

var providerSet = wire.NewSet(provider.NewPostgresqlClient, provider.NewRedisClient, provider.NewHttpClient, provider.NewEmailClient, provider.NewFilesystem, provider.NewRequestClient, wire.Struct(new(provider.Providers), "*"), cache.ProviderSet, logic.ProviderSet)
