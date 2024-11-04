package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(MessageUseCase), "*"),
	wire.Bind(new(MessageSendUseCase), new(*MessageUseCase)),

	NewAuthUseCase,
	NewContactUseCase,
	NewContactRequestUseCase,
	NewContactFolderUseCase,
	NewChatUseCase,
	NewGroupChatUseCase,
	NewGroupMemberUseCase,
	NewGroupChatAdsUseCase,
	NewGroupRequestUseCase,
	NewSplitUseCase,
	NewIpAddressUseCase,
	NewStickerUseCase,
	NewBotUseCase,
	NewUserUseCase,
)
