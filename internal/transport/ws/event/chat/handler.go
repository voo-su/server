package chat

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"voo.su/internal/usecase"
	"voo.su/pkg/socket"
)

type handle func(ctx context.Context, client socket.IClient, data []byte)

type Handler struct {
	Redis          *redis.Client
	MemberUseCase  *usecase.GroupChatMemberUseCase
	Handlers       map[string]func(ctx context.Context, client socket.IClient, data []byte)
	MessageUseCase usecase.IMessageUseCase
}

func NewHandler(
	redis *redis.Client,
	memberUseCase *usecase.GroupChatMemberUseCase,
) *Handler {
	return &Handler{
		Redis:         redis,
		MemberUseCase: memberUseCase,
	}
}

func (h *Handler) init() {
	h.Handlers = make(map[string]func(ctx context.Context, client socket.IClient, data []byte))
	h.Handlers["voo.message.publish"] = h.onPublish
	h.Handlers["voo.message.revoke"] = h.onRevokeMessage
	h.Handlers["voo.message.delete"] = h.onDeleteMessage
	h.Handlers["voo.message.read"] = h.onReadMessage
	h.Handlers["voo.message.keyboard"] = h.onKeyboardMessage
}

func (h *Handler) Call(ctx context.Context, client socket.IClient, event string, data []byte) {
	if h.Handlers == nil {
		h.init()
	}
	if call, ok := h.Handlers[event]; ok {
		call(ctx, client, data)
	} else {
		log.Printf("Событие чата: %s не зарегистрировано обратное вызов\n", event)
	}
}
