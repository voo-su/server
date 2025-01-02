// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

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
	NewFileSplitUseCase,
	NewIpAddressUseCase,
	NewStickerUseCase,
	NewBotUseCase,
	NewUserUseCase,
	NewProjectUseCase,
)
