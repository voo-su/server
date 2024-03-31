package chat

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"voo.su/internal/service"
	"voo.su/pkg/core/socket"
)

type handle func(ctx context.Context, client socket.IClient, data []byte)

type Handler struct {
	redis         *redis.Client
	memberService *service.GroupChatMemberService
	handlers      map[string]func(ctx context.Context, client socket.IClient, data []byte)
	message       service.MessageSendService
}

func NewHandler(
	redis *redis.Client,
	memberService *service.GroupChatMemberService,
	message *service.MessageService,
) *Handler {
	return &Handler{redis: redis, memberService: memberService, message: message}
}

func (h *Handler) init() {
	h.handlers = make(map[string]func(ctx context.Context, client socket.IClient, data []byte))
	h.handlers["voo.message.publish"] = h.onPublish
	h.handlers["voo.message.revoke"] = h.onRevokeMessage
	h.handlers["voo.message.delete"] = h.onDeleteMessage
	h.handlers["voo.message.read"] = h.onReadMessage
	h.handlers["voo.message.keyboard"] = h.onKeyboardMessage
}

func (h *Handler) Call(ctx context.Context, client socket.IClient, event string, data []byte) {
	if h.handlers == nil {
		h.init()
	}
	if call, ok := h.handlers[event]; ok {
		call(ctx, client, data)
	} else {
		log.Printf("Событие чата: %s не зарегистрировано обратное вызов\n", event)
	}
}
