// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"voo.su/internal/cli"
	"voo.su/internal/cli/handler/cron"
	"voo.su/internal/cli/handler/queue"
	"voo.su/internal/config"
	"voo.su/internal/delivery/grpc"
	handler3 "voo.su/internal/delivery/grpc/handler"
	"voo.su/internal/delivery/grpc/middleware"
	"voo.su/internal/delivery/http"
	"voo.su/internal/delivery/http/handler"
	"voo.su/internal/delivery/http/handler/bot"
	"voo.su/internal/delivery/http/handler/manager"
	"voo.su/internal/delivery/http/handler/v1"
	"voo.su/internal/delivery/http/router"
	"voo.su/internal/delivery/ws"
	"voo.su/internal/delivery/ws/consume"
	chat2 "voo.su/internal/delivery/ws/consume/chat"
	"voo.su/internal/delivery/ws/event"
	"voo.su/internal/delivery/ws/event/chat"
	handler2 "voo.su/internal/delivery/ws/handler"
	"voo.su/internal/delivery/ws/process"
	router2 "voo.su/internal/delivery/ws/router"
	"voo.su/internal/domain/logic"
	"voo.su/internal/infrastructure"
	repository3 "voo.su/internal/infrastructure/clickhouse/repository"
	repository2 "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/provider"
	"voo.su/internal/usecase"
)

// Injectors from wire.go:

func NewHttpInjector(conf *config.Config) *http.AppProvider {
	iLocale := provider.NewLocale(conf)
	email := provider.NewEmailClient(conf)
	client := provider.NewRedisClient(conf, iLocale)
	smsCacheRepository := repository.NewSmsCacheRepository(client)
	db := provider.NewPostgresqlClient(conf, iLocale)
	relationCacheRepository := repository.NewRelationCacheRepository(client)
	groupChatMemberRepository := repository2.NewGroupMemberRepository(db, relationCacheRepository)
	userRepository := repository2.NewUserRepository(db)
	userSessionRepository := repository2.NewUserSessionRepository(db)
	conn := provider.NewClickHouseClient(conf, iLocale)
	authCodeRepository := repository3.NewAuthCodeRepository(conn)
	jwtTokenCacheRepository := repository.NewJwtTokenCacheRepository(client)
	authUseCase := usecase.NewAuthUseCase(conf, iLocale, email, smsCacheRepository, groupChatMemberRepository, userRepository, userSessionRepository, authCodeRepository, jwtTokenCacheRepository)
	source := infrastructure.NewSource(db, conn, client)
	httpClient := provider.NewHttpClient()
	ipAddressUseCase := usecase.NewIpAddressUseCase(conf, iLocale, source, httpClient)
	chatRepository := repository2.NewChatRepository(db)
	redisLockCacheRepository := repository.NewRedisLockCacheRepository(client)
	serverCacheRepository := repository.NewServerCacheRepository(client)
	clientCacheRepository := repository.NewClientCacheRepository(conf, client, serverCacheRepository)
	messageCacheRepository := repository.NewMessageCacheRepository(client)
	unreadCacheRepository := repository.NewUnreadCacheRepository(client)
	chatUseCase := usecase.NewChatUseCase(iLocale, source, chatRepository, groupChatMemberRepository, redisLockCacheRepository, clientCacheRepository, messageCacheRepository, unreadCacheRepository)
	botRepository := repository2.NewBotRepository(db)
	iMinio := provider.NewMinioClient(conf)
	botUseCase := usecase.NewBotUseCase(conf, iLocale, source, botRepository, userRepository, iMinio)
	pushTokenRepository := repository2.NewPushTokenRepository(db)
	userUseCase := usecase.NewUserUseCase(iLocale, source, userRepository, userSessionRepository, pushTokenRepository)
	sequenceCacheRepository := repository.NewSequenceCacheRepository(client)
	sequenceRepository := repository2.NewSequenceRepository(db, sequenceCacheRepository)
	messageForward := logic.NewMessageForward(iLocale, source, sequenceRepository)
	fileSplitRepository := repository2.NewFileSplitRepository(db)
	voteCacheRepository := repository.NewVoteCacheRepository(client)
	messageVoteRepository := repository2.NewMessageVoteRepository(db, voteCacheRepository)
	messageRepository := repository2.NewMessageRepository(db)
	groupChatRepository := repository2.NewGroupChatRepository(db)
	contactRepository := repository2.NewContactRepository(db, relationCacheRepository)
	iNatsClient := provider.NewNatsClient(conf)
	messageUseCase := &usecase.MessageUseCase{
		Conf:                conf,
		Locale:              iLocale,
		Source:              source,
		MessageForward:      messageForward,
		Minio:               iMinio,
		GroupChatMemberRepo: groupChatMemberRepository,
		FileSplitRepo:       fileSplitRepository,
		MessageVoteRepo:     messageVoteRepository,
		Sequence:            sequenceRepository,
		MessageRepo:         messageRepository,
		BotRepo:             botRepository,
		GroupChatRepo:       groupChatRepository,
		ContactRepo:         contactRepository,
		UnreadCache:         unreadCacheRepository,
		MessageCache:        messageCacheRepository,
		ServerCache:         serverCacheRepository,
		ClientCache:         clientCacheRepository,
		Nats:                iNatsClient,
	}
	auth := &v1.Auth{
		Conf:             conf,
		Locale:           iLocale,
		AuthUseCase:      authUseCase,
		IpAddressUseCase: ipAddressUseCase,
		ChatUseCase:      chatUseCase,
		BotUseCase:       botUseCase,
		UserUseCase:      userUseCase,
		MessageUseCase:   messageUseCase,
	}
	account := &v1.Account{
		Locale:      iLocale,
		UserUseCase: userUseCase,
	}
	contactUseCase := usecase.NewContactUseCase(iLocale, source, contactRepository)
	contact := &v1.Contact{
		Locale:         iLocale,
		ContactUseCase: contactUseCase,
		ChatUseCase:    chatUseCase,
		UserUseCase:    userUseCase,
		MessageUseCase: messageUseCase,
	}
	contactRequestUseCase := usecase.NewContactRequestUseCase(iLocale, source)
	contactRequest := &v1.ContactRequest{
		Locale:                iLocale,
		ContactRequestUseCase: contactRequestUseCase,
		ContactUseCase:        contactUseCase,
		MessageUseCase:        messageUseCase,
	}
	groupChatUseCase := usecase.NewGroupChatUseCase(iLocale, source, groupChatRepository, groupChatMemberRepository, sequenceRepository, relationCacheRepository, redisLockCacheRepository)
	chat := &v1.Chat{
		Locale:           iLocale,
		ChatUseCase:      chatUseCase,
		MessageUseCase:   messageUseCase,
		ContactUseCase:   contactUseCase,
		GroupChatUseCase: groupChatUseCase,
		UserUseCase:      userUseCase,
	}
	groupChatMemberUseCase := usecase.NewGroupMemberUseCase(iLocale, source, groupChatMemberRepository)
	storageUseCase := usecase.NewStorageUseCase(iLocale, iMinio)
	message := &v1.Message{
		Conf:                   conf,
		Locale:                 iLocale,
		ChatUseCase:            chatUseCase,
		MessageUseCase:         messageUseCase,
		GroupChatMemberUseCase: groupChatMemberUseCase,
		StorageUseCase:         storageUseCase,
	}
	fileSplitUseCase := usecase.NewFileSplitUseCase(conf, iLocale, source, fileSplitRepository, iMinio)
	upload := &v1.Upload{
		Conf:             conf,
		Locale:           iLocale,
		Minio:            iMinio,
		FileSplitUseCase: fileSplitUseCase,
	}
	groupChat := &v1.GroupChat{
		Locale:                 iLocale,
		GroupChatUseCase:       groupChatUseCase,
		GroupChatMemberUseCase: groupChatMemberUseCase,
		ChatUseCase:            chatUseCase,
		ContactUseCase:         contactUseCase,
		MessageUseCase:         messageUseCase,
		UserUseCase:            userUseCase,
	}
	groupChatRequestRepository := repository2.NewGroupChatApplyRepository(db)
	groupChatRequestCacheRepository := repository.NewGroupChatRequestCacheRepository(client)
	groupChatRequestUseCase := usecase.NewGroupRequestUseCase(iLocale, source, groupChatRequestRepository, groupChatRequestCacheRepository)
	groupChatRequest := &v1.GroupChatRequest{
		Locale:                  iLocale,
		GroupChatRequestUseCase: groupChatRequestUseCase,
		GroupChatMemberUseCase:  groupChatMemberUseCase,
		GroupChatUseCase:        groupChatUseCase,
	}
	stickerRepository := repository2.NewStickerRepository(db)
	stickerUseCase := usecase.NewStickerUseCase(iLocale, source, stickerRepository, iMinio, redisLockCacheRepository)
	sticker := &v1.Sticker{
		Conf:           conf,
		Locale:         iLocale,
		StickerUseCase: stickerUseCase,
		StorageUseCase: storageUseCase,
	}
	contactFolderRepository := repository2.NewContactFolderRepository(db)
	contactFolderUseCase := usecase.NewContactFolderUseCase(iLocale, source, contactRepository, contactFolderRepository)
	contactFolder := &v1.ContactFolder{
		Locale:               iLocale,
		ContactFolderUseCase: contactFolderUseCase,
	}
	groupChatAdsRepository := repository2.NewGroupChatAdsRepository(db)
	groupChatAdsUseCase := usecase.NewGroupChatAdsUseCase(iLocale, source, groupChatAdsRepository)
	groupChatAds := &v1.GroupChatAds{
		Locale:              iLocale,
		GroupMemberUseCase:  groupChatMemberUseCase,
		GroupChatAdsUseCase: groupChatAdsUseCase,
		MessageUseCase:      messageUseCase,
	}
	search := &v1.Search{
		Locale:           iLocale,
		UserUseCase:      userUseCase,
		GroupChatUseCase: groupChatUseCase,
	}
	v1Bot := &v1.Bot{
		Locale:         iLocale,
		BotUseCase:     botUseCase,
		MessageUseCase: messageUseCase,
	}
	projectRepository := repository2.NewProjectRepository(db)
	projectMemberRepository := repository2.NewProjectMemberRepository(db)
	projectTaskTypeRepository := repository2.NewProjectTaskTypeRepository(db)
	projectTaskRepository := repository2.NewProjectTaskRepository(db)
	projectTaskCommentRepository := repository2.NewProjectTaskCommentRepository(db)
	projectContactCacheRepository := repository.NewProjectContactCacheRepository(client)
	projectTaskCoexecutorRepository := repository2.NewProjectTaskCoexecutorRepository(db)
	projectTaskWatcherRepository := repository2.NewProjectTaskWatcherRepository(db)
	projectUseCase := usecase.NewProjectUseCase(iLocale, source, projectRepository, projectMemberRepository, projectTaskTypeRepository, projectTaskRepository, projectTaskCommentRepository, userRepository, projectContactCacheRepository, projectTaskCoexecutorRepository, projectTaskWatcherRepository, redisLockCacheRepository)
	project := &v1.Project{
		Locale:         iLocale,
		ProjectUseCase: projectUseCase,
		ContactUseCase: contactUseCase,
	}
	projectTask := &v1.ProjectTask{
		Locale:         iLocale,
		ProjectUseCase: projectUseCase,
	}
	projectTaskComment := &v1.ProjectTaskComment{
		Locale:         iLocale,
		ProjectUseCase: projectUseCase,
	}
	v1Handler := &v1.Handler{
		Auth:               auth,
		Account:            account,
		Contact:            contact,
		ContactRequest:     contactRequest,
		Chat:               chat,
		Message:            message,
		Upload:             upload,
		GroupChat:          groupChat,
		GroupChatRequest:   groupChatRequest,
		Sticker:            sticker,
		ContactFolder:      contactFolder,
		GroupChatAds:       groupChatAds,
		Search:             search,
		Bot:                v1Bot,
		Project:            project,
		ProjectTask:        projectTask,
		ProjectTaskComment: projectTaskComment,
	}
	botMessage := &bot.Message{
		Locale:         iLocale,
		MessageUseCase: messageUseCase,
		BotUseCase:     botUseCase,
	}
	botHandler := &bot.Handler{
		Message: botMessage,
	}
	dashboard := &manager.Dashboard{
		Locale:           iLocale,
		UserUseCase:      userUseCase,
		MessageUseCase:   messageUseCase,
		GroupChatUseCase: groupChatUseCase,
	}
	managerHandler := &manager.Handler{
		Dashboard: dashboard,
	}
	handlerHandler := &handler.Handler{
		V1:      v1Handler,
		Bot:     botHandler,
		Manager: managerHandler,
	}
	accessLogRepository := repository3.NewAccessLogRepository(conn)
	engine := router.NewRouter(conf, iLocale, handlerHandler, jwtTokenCacheRepository, accessLogRepository)
	appProvider := &http.AppProvider{
		Conf:   conf,
		Engine: engine,
	}
	return appProvider
}

func NewWsInjector(conf *config.Config) *ws.AppProvider {
	iLocale := provider.NewLocale(conf)
	client := provider.NewRedisClient(conf, iLocale)
	serverCacheRepository := repository.NewServerCacheRepository(client)
	clientCacheRepository := repository.NewClientCacheRepository(conf, client, serverCacheRepository)
	roomCacheRepository := repository.NewRoomCacheRepository(client)
	db := provider.NewPostgresqlClient(conf, iLocale)
	relationCacheRepository := repository.NewRelationCacheRepository(client)
	groupChatMemberRepository := repository2.NewGroupMemberRepository(db, relationCacheRepository)
	conn := provider.NewClickHouseClient(conf, iLocale)
	source := infrastructure.NewSource(db, conn, client)
	groupChatMemberUseCase := usecase.NewGroupMemberUseCase(iLocale, source, groupChatMemberRepository)
	chatHandler := chat.NewHandler(client, groupChatMemberUseCase)
	chatEvent := &event.ChatEvent{
		Redis:                  client,
		Conf:                   conf,
		RoomCache:              roomCacheRepository,
		GroupChatMemberRepo:    groupChatMemberRepository,
		GroupChatMemberUseCase: groupChatMemberUseCase,
		Handler:                chatHandler,
	}
	chatChannel := &handler2.ChatChannel{
		ClientCacheRepo: clientCacheRepository,
		Event:           chatEvent,
	}
	handlerHandler := &handler2.Handler{
		Chat: chatChannel,
		Conf: conf,
	}
	jwtTokenCacheRepository := repository.NewJwtTokenCacheRepository(client)
	accessLogRepository := repository3.NewAccessLogRepository(conn)
	engine := router2.NewRouter(conf, iLocale, handlerHandler, jwtTokenCacheRepository, accessLogRepository)
	healthSubscribe := process.NewHealthSubscribe(conf, serverCacheRepository)
	chatRepository := repository2.NewChatRepository(db)
	redisLockCacheRepository := repository.NewRedisLockCacheRepository(client)
	messageCacheRepository := repository.NewMessageCacheRepository(client)
	unreadCacheRepository := repository.NewUnreadCacheRepository(client)
	chatUseCase := usecase.NewChatUseCase(iLocale, source, chatRepository, groupChatMemberRepository, redisLockCacheRepository, clientCacheRepository, messageCacheRepository, unreadCacheRepository)
	sequenceCacheRepository := repository.NewSequenceCacheRepository(client)
	sequenceRepository := repository2.NewSequenceRepository(db, sequenceCacheRepository)
	messageForward := logic.NewMessageForward(iLocale, source, sequenceRepository)
	iMinio := provider.NewMinioClient(conf)
	fileSplitRepository := repository2.NewFileSplitRepository(db)
	voteCacheRepository := repository.NewVoteCacheRepository(client)
	messageVoteRepository := repository2.NewMessageVoteRepository(db, voteCacheRepository)
	messageRepository := repository2.NewMessageRepository(db)
	botRepository := repository2.NewBotRepository(db)
	groupChatRepository := repository2.NewGroupChatRepository(db)
	contactRepository := repository2.NewContactRepository(db, relationCacheRepository)
	iNatsClient := provider.NewNatsClient(conf)
	messageUseCase := &usecase.MessageUseCase{
		Conf:                conf,
		Locale:              iLocale,
		Source:              source,
		MessageForward:      messageForward,
		Minio:               iMinio,
		GroupChatMemberRepo: groupChatMemberRepository,
		FileSplitRepo:       fileSplitRepository,
		MessageVoteRepo:     messageVoteRepository,
		Sequence:            sequenceRepository,
		MessageRepo:         messageRepository,
		BotRepo:             botRepository,
		GroupChatRepo:       groupChatRepository,
		ContactRepo:         contactRepository,
		UnreadCache:         unreadCacheRepository,
		MessageCache:        messageCacheRepository,
		ServerCache:         serverCacheRepository,
		ClientCache:         clientCacheRepository,
		Nats:                iNatsClient,
	}
	contactUseCase := usecase.NewContactUseCase(iLocale, source, contactRepository)
	handler3 := &chat2.Handler{
		Conf:           conf,
		Locale:         iLocale,
		Source:         source,
		ClientCache:    clientCacheRepository,
		RoomCache:      roomCacheRepository,
		ChatUseCase:    chatUseCase,
		MessageUseCase: messageUseCase,
		ContactUseCase: contactUseCase,
	}
	chatSubscribe := consume.NewChatSubscribe(handler3)
	messageSubscribe := process.NewMessageSubscribe(conf, client, chatSubscribe)
	subServers := &process.SubServers{
		HealthSubscribe:  healthSubscribe,
		MessageSubscribe: messageSubscribe,
	}
	server := process.NewServer(subServers)
	email := provider.NewEmailClient(conf)
	providers := &provider.Providers{
		EmailClient: email,
	}
	appProvider := &ws.AppProvider{
		Conf:      conf,
		Locale:    iLocale,
		Engine:    engine,
		Coroutine: server,
		Handler:   handlerHandler,
		Providers: providers,
	}
	return appProvider
}

func NewGrpcInjector(conf *config.Config) *grpc.AppProvider {
	iLocale := provider.NewLocale(conf)
	authMiddleware := middleware.NewAuthMiddleware(conf, iLocale)
	grpcMethodService := middleware.NewGrpMethodsService()
	email := provider.NewEmailClient(conf)
	client := provider.NewRedisClient(conf, iLocale)
	smsCacheRepository := repository.NewSmsCacheRepository(client)
	db := provider.NewPostgresqlClient(conf, iLocale)
	relationCacheRepository := repository.NewRelationCacheRepository(client)
	groupChatMemberRepository := repository2.NewGroupMemberRepository(db, relationCacheRepository)
	userRepository := repository2.NewUserRepository(db)
	userSessionRepository := repository2.NewUserSessionRepository(db)
	conn := provider.NewClickHouseClient(conf, iLocale)
	authCodeRepository := repository3.NewAuthCodeRepository(conn)
	jwtTokenCacheRepository := repository.NewJwtTokenCacheRepository(client)
	authUseCase := usecase.NewAuthUseCase(conf, iLocale, email, smsCacheRepository, groupChatMemberRepository, userRepository, userSessionRepository, authCodeRepository, jwtTokenCacheRepository)
	source := infrastructure.NewSource(db, conn, client)
	httpClient := provider.NewHttpClient()
	ipAddressUseCase := usecase.NewIpAddressUseCase(conf, iLocale, source, httpClient)
	chatRepository := repository2.NewChatRepository(db)
	redisLockCacheRepository := repository.NewRedisLockCacheRepository(client)
	serverCacheRepository := repository.NewServerCacheRepository(client)
	clientCacheRepository := repository.NewClientCacheRepository(conf, client, serverCacheRepository)
	messageCacheRepository := repository.NewMessageCacheRepository(client)
	unreadCacheRepository := repository.NewUnreadCacheRepository(client)
	chatUseCase := usecase.NewChatUseCase(iLocale, source, chatRepository, groupChatMemberRepository, redisLockCacheRepository, clientCacheRepository, messageCacheRepository, unreadCacheRepository)
	botRepository := repository2.NewBotRepository(db)
	iMinio := provider.NewMinioClient(conf)
	botUseCase := usecase.NewBotUseCase(conf, iLocale, source, botRepository, userRepository, iMinio)
	sequenceCacheRepository := repository.NewSequenceCacheRepository(client)
	sequenceRepository := repository2.NewSequenceRepository(db, sequenceCacheRepository)
	messageForward := logic.NewMessageForward(iLocale, source, sequenceRepository)
	fileSplitRepository := repository2.NewFileSplitRepository(db)
	voteCacheRepository := repository.NewVoteCacheRepository(client)
	messageVoteRepository := repository2.NewMessageVoteRepository(db, voteCacheRepository)
	messageRepository := repository2.NewMessageRepository(db)
	groupChatRepository := repository2.NewGroupChatRepository(db)
	contactRepository := repository2.NewContactRepository(db, relationCacheRepository)
	iNatsClient := provider.NewNatsClient(conf)
	messageUseCase := &usecase.MessageUseCase{
		Conf:                conf,
		Locale:              iLocale,
		Source:              source,
		MessageForward:      messageForward,
		Minio:               iMinio,
		GroupChatMemberRepo: groupChatMemberRepository,
		FileSplitRepo:       fileSplitRepository,
		MessageVoteRepo:     messageVoteRepository,
		Sequence:            sequenceRepository,
		MessageRepo:         messageRepository,
		BotRepo:             botRepository,
		GroupChatRepo:       groupChatRepository,
		ContactRepo:         contactRepository,
		UnreadCache:         unreadCacheRepository,
		MessageCache:        messageCacheRepository,
		ServerCache:         serverCacheRepository,
		ClientCache:         clientCacheRepository,
		Nats:                iNatsClient,
	}
	auth := handler3.NewAuthHandler(conf, iLocale, authUseCase, ipAddressUseCase, chatUseCase, botUseCase, messageUseCase)
	pushTokenRepository := repository2.NewPushTokenRepository(db)
	userUseCase := usecase.NewUserUseCase(iLocale, source, userRepository, userSessionRepository, pushTokenRepository)
	account := handler3.NewAccountHandler(userUseCase)
	contactUseCase := usecase.NewContactUseCase(iLocale, source, contactRepository)
	handlerChat := handler3.NewChatHandler(conf, iLocale, contactUseCase, chatUseCase)
	message := handler3.NewMessageHandler(conf, iLocale, messageUseCase)
	contact := handler3.NewContactHandler(conf, iLocale, contactUseCase)
	appProvider := &grpc.AppProvider{
		Conf:           conf,
		AuthMiddleware: authMiddleware,
		RoutesServices: grpcMethodService,
		AuthHandler:    auth,
		AccountHandler: account,
		ChatHandler:    handlerChat,
		MessageHandler: message,
		ContactHandler: contact,
	}
	return appProvider
}

func NewCronInjector(conf *config.Config) *cli.CronProvider {
	iLocale := provider.NewLocale(conf)
	client := provider.NewRedisClient(conf, iLocale)
	serverCacheRepository := repository.NewServerCacheRepository(client)
	clearWsCache := cron.NewClearWsCache(serverCacheRepository)
	db := provider.NewPostgresqlClient(conf, iLocale)
	iMinio := provider.NewMinioClient(conf)
	clearTmpFile := cron.NewClearTmpFile(conf, db, iMinio)
	clearExpireServer := cron.NewClearExpireServer(serverCacheRepository)
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
	iLocale := provider.NewLocale(conf)
	db := provider.NewPostgresqlClient(conf, iLocale)
	client := provider.NewRedisClient(conf, iLocale)
	emailHandle := queue.EmailHandle{
		Redis: client,
	}
	iNatsClient := provider.NewNatsClient(conf)
	pushHandle := queue.PushHandle{
		Conf: conf,
		Nats: iNatsClient,
	}
	queueJobs := &cli.QueueJobs{
		EmailHandle: emailHandle,
		PushHandle:  pushHandle,
	}
	queueProvider := &cli.QueueProvider{
		Conf: conf,
		DB:   db,
		Jobs: queueJobs,
	}
	return queueProvider
}

func NewMigrateInjector(conf *config.Config) *cli.MigrateProvider {
	iLocale := provider.NewLocale(conf)
	db := provider.NewPostgresqlClient(conf, iLocale)
	migrateProvider := &cli.MigrateProvider{
		Conf: conf,
		DB:   db,
	}
	return migrateProvider
}

func NewGenerateInjector(conf *config.Config) *cli.GenerateProvider {
	generateProvider := &cli.GenerateProvider{}
	return generateProvider
}
