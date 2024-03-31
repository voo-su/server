package consume

import (
	"github.com/google/wire"
	"voo.su/internal/transport/ws/consume/chat"
)

var ProviderSet = wire.NewSet(
	NewChatSubscribe,
	chat.NewHandler,
)
