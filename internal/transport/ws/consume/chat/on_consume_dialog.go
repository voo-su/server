package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/logger"
)

type ConsumeDialog struct {
	DialogType int   `json:"dialog_type"`
	SenderID   int64 `json:"sender_id"`
	ReceiverID int64 `json:"receiver_id"`
	RecordID   int64 `json:"record_id"`
}

func (h *Handler) onConsumeDialog(ctx context.Context, body []byte) {
	var in ConsumeDialog
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeDialog Ошибка при декодировании: ", err.Error())
		return
	}

	var clientIds []int64
	if in.DialogType == entity.ChatPrivateMode {
		for _, val := range [2]int64{in.SenderID, in.ReceiverID} {
			ids := h.ClientStorage.GetUidFromClientIds(ctx, h.Config.ServerId(), socket.Session.Chat.Name(), strconv.FormatInt(val, 10))

			clientIds = append(clientIds, ids...)
		}
	} else if in.DialogType == entity.ChatGroupMode {
		ids := h.RoomStorage.All(ctx, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: entity.RoomImGroup,
			Number:   strconv.Itoa(int(in.ReceiverID)),
			Sid:      h.Config.ServerId(),
		})
		clientIds = append(clientIds, ids...)
	}
	if len(clientIds) == 0 {
		return
	}

	data, err := h.MessageService.GetDialogRecord(ctx, in.RecordID)
	if err != nil {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetAck(true)
	c.SetMessage(entity.PushEventImMessage, map[string]any{
		"sender_id":   in.SenderID,
		"receiver_id": in.ReceiverID,
		"dialog_type": in.DialogType,
		"data":        data,
	})
	socket.Session.Chat.Write(c)
}
