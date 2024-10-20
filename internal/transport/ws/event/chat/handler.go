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
	Redis         *redis.Client
	MemberService *service.GroupChatMemberService
	Handlers      map[string]func(ctx context.Context, client socket.IClient, data []byte)
	Message       service.MessageSendService
}

func NewHandler(
	redis *redis.Client,
	memberService *service.GroupChatMemberService,
	message *service.MessageService,
) *Handler {
	return &Handler{
		Redis:         redis,
		MemberService: memberService,
		Message:       message,
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
