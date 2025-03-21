package chat

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/pkg/socket"
)

func (h *Handler) onConsumeChatKeyboard(ctx context.Context, body []byte) {
	var in entity.ConsumeChatKeyboard
	if err := json.Unmarshal(body, &in); err != nil {
		log.Fatalf("onConsumeChatKeyboard json decode err: %s", err.Error())
		return
	}

	ids := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(in.ReceiverId))
	if len(ids) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(ids...)
	c.SetMessage(constant.PushEventImMessageKeyboard, map[string]any{
		"sender_id":   in.SenderId,
		"receiver_id": in.ReceiverId,
	})
	socket.Session.Chat.Write(c)
}
