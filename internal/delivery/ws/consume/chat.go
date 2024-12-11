package consume

import (
	"context"
	"voo.su/internal/delivery/ws/consume/chat"
)

type ChatSubscribe struct {
	Handler *chat.Handler
}

func NewChatSubscribe(handel *chat.Handler) *ChatSubscribe {
	return &ChatSubscribe{
		Handler: handel,
	}
}

func (c *ChatSubscribe) Call(event string, data []byte) {
	c.Handler.Call(context.TODO(), event, data)
}
