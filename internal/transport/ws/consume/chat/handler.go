package chat

import (
	"context"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/repository"
	"voo.su/internal/repository/cache"
	"voo.su/internal/usecase"
)

var handlers map[string]func(ctx context.Context, data []byte)

type Handler struct {
	Conf           *config.Config
	ClientCache    *cache.ClientCache
	RoomCache      *cache.RoomCache
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase
	ContactUseCase *usecase.ContactUseCase
	Source         *repository.Source
}

func NewHandler(
	conf *config.Config,
	clientCache *cache.ClientCache,
	roomCache *cache.RoomCache,
	chatUseCase *usecase.ChatUseCase,
	contactUseCase *usecase.ContactUseCase,
	source *repository.Source,
) *Handler {
	return &Handler{
		Conf:           conf,
		ClientCache:    clientCache,
		RoomCache:      roomCache,
		ChatUseCase:    chatUseCase,
		ContactUseCase: contactUseCase,
		Source:         source,
	}
}

func (h *Handler) init() {
	handlers = make(map[string]func(ctx context.Context, data []byte))
	handlers[constant.SubEventImMessage] = h.onConsumeDialog
	handlers[constant.SubEventImMessageKeyboard] = h.onConsumeDialogKeyboard
	handlers[constant.SubEventImMessageRead] = h.onConsumeDialogRead
	handlers[constant.SubEventImMessageRevoke] = h.onConsumeDialogRevoke
	handlers[constant.SubEventContactStatus] = h.onConsumeContactStatus
	handlers[constant.SubEventContactRequest] = h.onConsumeContactApply
	handlers[constant.SubEventGroupChatJoin] = h.onConsumeGroupJoin
	handlers[constant.SubEventGroupChatRequest] = h.onConsumeGroupApply
}

func (h *Handler) Call(ctx context.Context, event string, data []byte) {
	if handlers == nil {
		h.init()
	}
	if call, ok := handlers[event]; ok {
		call(ctx, data)
	} else {
		log.Printf("Незарегистрированное событие обратного вызова %s \n", event)
	}
}
