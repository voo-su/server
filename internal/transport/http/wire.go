package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"voo.su/internal/config"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/internal/transport/http/handler"
	"voo.su/internal/transport/http/handler/bot"
	v1 "voo.su/internal/transport/http/handler/v1"
	"voo.su/internal/transport/http/router"
)

type AppProvider struct {
	Conf   *config.Config
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	wire.Struct(new(handler.Handler), "*"),

	router.NewRouter,

	// v1
	wire.Struct(new(v1.Handler), "*"),
	wire.Struct(new(v1.Auth), "*"),
	wire.Struct(new(v1.Account), "*"),
	wire.Struct(new(v1.Upload), "*"),
	wire.Struct(new(v1.Contact), "*"),
	wire.Struct(new(v1.ContactRequest), "*"),
	wire.Struct(new(v1.Chat), "*"),
	wire.Struct(new(v1.Message), "*"),
	wire.Struct(new(v1.Publish), "*"),
	wire.Struct(new(v1.GroupChat), "*"),
	wire.Struct(new(v1.GroupChatRequest), "*"),
	wire.Struct(new(v1.Sticker), "*"),
	wire.Struct(new(v1.ContactFolder), "*"),
	wire.Struct(new(v1.GroupChatAds), "*"),
	wire.Struct(new(v1.Search), "*"),

	// Bot
	wire.Struct(new(bot.Handler), "*"),
	wire.Struct(new(bot.Message), "*"),

	wire.Struct(new(service.MessageService), "*"),
	wire.Bind(new(service.MessageSendService), new(*service.MessageService)),
	service.NewAuthService,
	service.NewContactService,
	service.NewContactApplyService,
	service.NewContactFolderService,
	service.NewChatService,
	service.NewGroupChatService,
	service.NewGroupMemberService,
	service.NewGroupChatAdsService,
	service.NewGroupRequestService,
	service.NewSplitService,
	service.NewIpAddressService,
	service.NewStickerService,

	repo.NewUserSession,
	repo.NewSource,
	repo.NewContact,
	repo.NewContactFolder,
	repo.NewGroupMember,
	repo.NewUser,
	repo.NewGroupChat,
	repo.NewGroupChatApply,
	repo.NewGroupChatAds,
	repo.NewDialog,
	repo.NewMessage,
	repo.NewMessageVote,
	repo.NewFileSplit,
	repo.NewSequence,
	repo.NewBot,
	repo.NewSticker,
)
