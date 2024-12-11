package chat

import (
	"context"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/infrastructure"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
)

var handlers map[string]func(ctx context.Context, data []byte)

type Handler struct {
	Source         *infrastructure.Source
	Conf           *config.Config
	ClientCache    *redisRepo.ClientCacheRepository
	RoomCache      *redisRepo.RoomCacheRepository
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase
	ContactUseCase *usecase.ContactUseCase
}

func NewHandler(
	source *infrastructure.Source,
	conf *config.Config,
	clientCache *redisRepo.ClientCacheRepository,
	roomCache *redisRepo.RoomCacheRepository,
	chatUseCase *usecase.ChatUseCase,
	contactUseCase *usecase.ContactUseCase,

) *Handler {
	return &Handler{
		Source:         source,
		Conf:           conf,
		ClientCache:    clientCache,
		RoomCache:      roomCache,
		ChatUseCase:    chatUseCase,
		ContactUseCase: contactUseCase,
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
