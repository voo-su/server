package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	"voo.su/pkg/logger"
	"voo.su/pkg/socket"
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
	if in.DialogType == constant.ChatPrivateMode {
		for _, val := range [2]int64{in.SenderID, in.ReceiverID} {
			ids := h.ClientStorage.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.FormatInt(val, 10))

			clientIds = append(clientIds, ids...)
		}
	} else if in.DialogType == constant.ChatGroupMode {
		ids := h.RoomStorage.All(ctx, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: constant.RoomImGroup,
			Number:   strconv.Itoa(int(in.ReceiverID)),
			Sid:      h.Conf.ServerId(),
		})
		clientIds = append(clientIds, ids...)
	}
	if len(clientIds) == 0 {
		return
	}

	data, err := h.MessageUseCase.GetDialogRecord(ctx, in.RecordID)
	if err != nil {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetAck(true)
	c.SetMessage(constant.PushEventImMessage, map[string]any{
		"sender_id":   in.SenderID,
		"receiver_id": in.ReceiverID,
		"dialog_type": in.DialogType,
		"data":        data,
	})
	socket.Session.Chat.Write(c)
}
