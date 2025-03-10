package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/constant"
	redisModel "voo.su/internal/infrastructure/redis/model"
	"voo.su/pkg/logger"
	"voo.su/pkg/socket"
)

type ConsumeMessage struct {
	ChatType   int   `json:"chat_type"`
	SenderId   int64 `json:"sender_id"`
	ReceiverId int64 `json:"receiver_id"`
	MessageId  int64 `json:"message_id"`
}

func (h *Handler) onConsumeMessage(ctx context.Context, body []byte) {
	var in ConsumeMessage
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeMessage json decode err: %s", err.Error())
		return
	}

	var clientIds []int64
	if in.ChatType == constant.ChatPrivateMode {
		for _, val := range [2]int64{in.SenderId, in.ReceiverId} {
			ids := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.FormatInt(val, 10))

			clientIds = append(clientIds, ids...)
		}
	} else if in.ChatType == constant.ChatGroupMode {
		ids := h.RoomCache.All(ctx, &redisModel.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: constant.RoomImGroup,
			Number:   strconv.Itoa(int(in.ReceiverId)),
			Sid:      h.Conf.ServerId(),
		})
		clientIds = append(clientIds, ids...)
	}
	if len(clientIds) == 0 {
		return
	}

	data, err := h.MessageUseCase.GetMessage(ctx, in.MessageId)
	if err != nil {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetAck(true)
	c.SetMessage(constant.PushEventImMessage, map[string]any{
		"sender_id":   in.SenderId,
		"receiver_id": in.ReceiverId,
		"chat_type":   in.ChatType,
		"data":        data,
	})
	socket.Session.Chat.Write(c)
}
