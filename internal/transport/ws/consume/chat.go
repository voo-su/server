package consume

import (
	"context"
	"voo.su/internal/transport/ws/consume/chat"
)

type ChatSubscribe struct {
	handler *chat.Handler
}

func NewChatSubscribe(handel *chat.Handler) *ChatSubscribe {
	return &ChatSubscribe{handler: handel}
}

func (s *ChatSubscribe) Call(event string, data []byte) {
	s.handler.Call(context.TODO(), event, data)
}
