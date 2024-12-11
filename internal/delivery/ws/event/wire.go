package event

import (
	"github.com/google/wire"
	"voo.su/internal/delivery/ws/event/chat"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(ChatEvent), "*"),
	chat.NewHandler,
)
