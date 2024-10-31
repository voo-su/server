package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	wire.Struct(new(MessageService), "*"),
	wire.Bind(new(MessageSendService), new(*MessageService)),
	NewAuthService,
	NewContactService,
	NewContactApplyService,
	NewContactFolderService,
	NewChatService,
	NewGroupChatService,
	NewGroupMemberService,
	NewGroupChatAdsService,
	NewGroupRequestService,
	NewSplitService,
	NewIpAddressService,
	NewStickerService,
	NewBotService,
)
