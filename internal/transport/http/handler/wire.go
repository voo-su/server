package handler

import (
	"github.com/google/wire"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
	"voo.su/internal/transport/http/handler/v1"
)

var ProviderSet = wire.NewSet(
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

	// v1
	wire.Struct(new(V1), "*"),
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
)
