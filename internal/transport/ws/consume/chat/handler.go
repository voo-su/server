package chat

import (
	"context"
	"log"
	"voo.su/internal/config"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/repo"
	"voo.su/internal/service"
)

var handlers map[string]func(ctx context.Context, data []byte)

type Handler struct {
	Conf           *config.Config
	ClientStorage  *cache.ClientStorage
	RoomStorage    *cache.RoomStorage
	DialogService  *service.DialogService
	MessageService *service.MessageService
	ContactService *service.ContactService
	Source         *repo.Source
}

func NewHandler(
	conf *config.Config,
	clientStorage *cache.ClientStorage,
	roomStorage *cache.RoomStorage,
	dialogService *service.DialogService,
	messageService *service.MessageService,
	contactService *service.ContactService,
	source *repo.Source,
) *Handler {
	return &Handler{
		Conf:           conf,
		ClientStorage:  clientStorage,
		RoomStorage:    roomStorage,
		DialogService:  dialogService,
		MessageService: messageService,
		ContactService: contactService,
		Source:         source,
	}
}

func (h *Handler) init() {
	handlers = make(map[string]func(ctx context.Context, data []byte))
	handlers[entity.SubEventImMessage] = h.onConsumeDialog
	handlers[entity.SubEventImMessageKeyboard] = h.onConsumeDialogKeyboard
	handlers[entity.SubEventImMessageRead] = h.onConsumeDialogRead
	handlers[entity.SubEventImMessageRevoke] = h.onConsumeDialogRevoke
	handlers[entity.SubEventContactStatus] = h.onConsumeContactStatus
	handlers[entity.SubEventContactRequest] = h.onConsumeContactApply
	handlers[entity.SubEventGroupChatJoin] = h.onConsumeGroupJoin
	handlers[entity.SubEventGroupChatRequest] = h.onConsumeGroupApply
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
