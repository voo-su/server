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

type ConsumeDialog struct {
	DialogType int   `json:"dialog_type"`
	SenderId   int64 `json:"sender_id"`
	ReceiverId int64 `json:"receiver_id"`
	RecordId   int64 `json:"record_id"`
}

func (h *Handler) onConsumeDialog(ctx context.Context, body []byte) {
	var in ConsumeDialog
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeDialog Ошибка при декодировании: %s", err.Error())
		return
	}

	var clientIds []int64
	if in.DialogType == constant.ChatPrivateMode {
		for _, val := range [2]int64{in.SenderId, in.ReceiverId} {
			ids := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.FormatInt(val, 10))

			clientIds = append(clientIds, ids...)
		}
	} else if in.DialogType == constant.ChatGroupMode {
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

	data, err := h.MessageUseCase.GetDialogRecord(ctx, in.RecordId)
	if err != nil {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetAck(true)
	c.SetMessage(constant.PushEventImMessage, map[string]any{
		"sender_id":   in.SenderId,
		"receiver_id": in.ReceiverId,
		"dialog_type": in.DialogType,
		"data":        data,
	})
	socket.Session.Chat.Write(c)
}
