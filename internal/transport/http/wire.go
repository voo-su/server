package http

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"voo.su/internal/config"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	handler2 "voo.su/internal/transport/http/handler"
	"voo.su/internal/transport/http/handler/v1"
	"voo.su/internal/transport/http/router"
	"voo.su/internal/transport/ws/handler"
)

type AppProvider struct {
	Config *config.Config
	Engine *gin.Engine
}

var ProviderSet = wire.NewSet(
	wire.Struct(new(AppProvider), "*"),
	wire.Struct(new(handler.Handler), "*"),

	router.NewRouter,

	// v1
	wire.Struct(new(handler2.V1), "*"),
	wire.Struct(new(v1.Auth), "*"),
	wire.Struct(new(v1.Account), "*"),
	wire.Struct(new(v1.User), "*"),
	wire.Struct(new(v1.Upload), "*"),
	wire.Struct(new(v1.Contact), "*"),
	wire.Struct(new(v1.ContactRequest), "*"),
	wire.Struct(new(v1.GroupChat), "*"),
	wire.Struct(new(v1.GroupChatRequest), "*"),
	wire.Struct(new(v1.Dialog), "*"),
	wire.Struct(new(v1.Message), "*"),
	wire.Struct(new(v1.Publish), "*"),
	wire.Struct(new(v1.Sticker), "*"),
	wire.Struct(new(v1.ContactGroup), "*"),
	wire.Struct(new(v1.GroupChatNotice), "*"),

	wire.Struct(new(service.MessageService), "*"),
	wire.Bind(new(service.MessageSendService), new(*service.MessageService)),
	service.NewAuthService,
	service.NewDialogService,
	service.NewContactService,
	service.NewContactApplyService,
	service.NewContactGroupService,
	service.NewGroupChatService,
	service.NewGroupMemberService,
	service.NewGroupChatNoticeService,
	service.NewGroupRequestService,
	service.NewSplitService,
	service.NewIpAddressService,
	service.NewStickerService,

	repo.NewUserSession,
	repo.NewSource,
	repo.NewContact,
	repo.NewContactGroup,
	repo.NewGroupMember,
	repo.NewUser,
	repo.NewGroupChat,
	repo.NewGroupChatApply,
	repo.NewGroupChatNotice,
	repo.NewDialog,
	repo.NewMessage,
	repo.NewMessageVote,
	repo.NewFileSplit,
	repo.NewSequence,
	repo.NewBot,
	repo.NewSticker,
)
