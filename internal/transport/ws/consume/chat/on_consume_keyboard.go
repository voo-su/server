package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/constant"
	"voo.su/pkg/logger"
	"voo.su/pkg/socket"
)

type ConsumeDialogKeyboard struct {
	SenderID   int `json:"sender_id"`
	ReceiverID int `json:"receiver_id"`
}

func (h *Handler) onConsumeDialogKeyboard(ctx context.Context, body []byte) {
	var in ConsumeDialogKeyboard
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeDialogKeyboard Ошибка при декодировании: ", err.Error())
		return
	}

	ids := h.ClientStorage.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(in.ReceiverID))
	if len(ids) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(ids...)
	c.SetMessage(constant.PushEventImMessageKeyboard, map[string]any{
		"sender_id":   in.SenderID,
		"receiver_id": in.ReceiverID,
	})
	socket.Session.Chat.Write(c)
}
