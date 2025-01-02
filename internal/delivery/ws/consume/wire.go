// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package consume

import (
	"github.com/google/wire"
	"voo.su/internal/delivery/ws/consume/chat"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(chat.Handler), "*"),
	NewChatSubscribe,
	//chat.NewHandler,
)
