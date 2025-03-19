package usecase

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(MessageUseCase), "*"),
	wire.Bind(new(IMessageUseCase), new(*MessageUseCase)),

	NewAuthUseCase,
	NewContactUseCase,
	NewContactRequestUseCase,
	NewContactFolderUseCase,
	NewChatUseCase,
	NewGroupChatUseCase,
	NewGroupMemberUseCase,
	NewGroupChatAdsUseCase,
	NewGroupRequestUseCase,
	NewUploadUseCase,
	NewIpAddressUseCase,
	NewStickerUseCase,
	NewBotUseCase,
	NewUserUseCase,
	NewProjectUseCase,
	NewStorageUseCase,
)
