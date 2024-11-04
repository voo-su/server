package chat

import (
	"context"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/repo"
	"voo.su/internal/usecase"
)

var handlers map[string]func(ctx context.Context, data []byte)

type Handler struct {
	Conf           *config.Config
	ClientStorage  *cache.ClientStorage
	RoomStorage    *cache.RoomStorage
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase *usecase.MessageUseCase
	ContactUseCase *usecase.ContactUseCase
	Source         *repo.Source
}

func NewHandler(
	conf *config.Config,
	clientStorage *cache.ClientStorage,
	roomStorage *cache.RoomStorage,
	chatUseCase *usecase.ChatUseCase,
	messageUseCase *usecase.MessageUseCase,
	contactUseCase *usecase.ContactUseCase,
	source *repo.Source,
) *Handler {
	return &Handler{
		Conf:           conf,
		ClientStorage:  clientStorage,
		RoomStorage:    roomStorage,
		ChatUseCase:    chatUseCase,
		MessageUseCase: messageUseCase,
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
