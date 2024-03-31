package event

import (
	"github.com/google/wire"
	"voo.su/internal/transport/ws/event/chat"
)

var ProviderSet = wire.NewSet(
	wire.Struct(new(ChatEvent), "*"),
	chat.NewHandler,
)
