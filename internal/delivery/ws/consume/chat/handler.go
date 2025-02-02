package chat

import (
	"context"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/constant"
	"voo.su/internal/infrastructure"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/internal/usecase"
	"voo.su/pkg/locale"
)

var handlers map[string]func(ctx context.Context, data []byte)

type Handler struct {
	Conf           *config.Config
	Locale         locale.ILocale
	Source         *infrastructure.Source
	ClientCache    *redisRepo.ClientCacheRepository
	RoomCache      *redisRepo.RoomCacheRepository
	ChatUseCase    *usecase.ChatUseCase
	MessageUseCase usecase.IMessageUseCase
	ContactUseCase *usecase.ContactUseCase
}

func (h *Handler) init() {
	handlers = make(map[string]func(ctx context.Context, data []byte))
	handlers[constant.SubEventImMessage] = h.onConsumeMessage
	handlers[constant.SubEventImMessageKeyboard] = h.onConsumeChatKeyboard
	handlers[constant.SubEventImMessageRead] = h.onConsumeMessageRead
	handlers[constant.SubEventImMessageRevoke] = h.onConsumeMessageRevoke
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
