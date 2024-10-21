package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/entity"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/logger"
)

type ConsumeDialogRead struct {
	SenderId   int      `json:"sender_id"`
	ReceiverId int      `json:"receiver_id"`
	MsgIds     []string `json:"msg_ids"`
}

func (h *Handler) onConsumeDialogRead(ctx context.Context, body []byte) {
	var in ConsumeDialogRead
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeContactApply Ошибка при декодировании: ", err.Error())
		return
	}

	clientIds := h.ClientStorage.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(in.ReceiverId))
	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(entity.PushEventImMessageRead, map[string]any{
		"sender_id":   in.SenderId,
		"receiver_id": in.ReceiverId,
		"msg_ids":     in.MsgIds,
	})
	socket.Session.Chat.Write(c)
}
