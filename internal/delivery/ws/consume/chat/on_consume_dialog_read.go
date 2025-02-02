package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/constant"
	"voo.su/pkg/logger"
	"voo.su/pkg/socket"
)

type ConsumeMessageRead struct {
	SenderId   int      `json:"sender_id"`
	ReceiverId int      `json:"receiver_id"`
	MsgIds     []string `json:"msg_ids"`
}

func (h *Handler) onConsumeMessageRead(ctx context.Context, body []byte) {
	var in ConsumeMessageRead
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeContactApply json decode err: %s", err.Error())
		return
	}

	clientIds := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(in.ReceiverId))
	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(constant.PushEventImMessageRead, map[string]any{
		"sender_id":   in.SenderId,
		"receiver_id": in.ReceiverId,
		"msg_ids":     in.MsgIds,
	})
	socket.Session.Chat.Write(c)
}
