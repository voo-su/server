// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package event

import (
	"github.com/google/wire"
	"voo.su/internal/delivery/ws/event/chat"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(ChatEvent), "*"),
	chat.NewHandler,
)
