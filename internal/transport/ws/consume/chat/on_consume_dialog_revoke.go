package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/entity"
	"voo.su/internal/repository/cache"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/logger"
)

type ConsumeDialogRevoke struct {
	MsgId string `json:"msg_id"`
}

func (h *Handler) onConsumeDialogRevoke(ctx context.Context, body []byte) {
	var in ConsumeDialogRevoke
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeDialogRevoke Ошибка при декодировании: ", err.Error())
		return
	}

	var record model.Message
	if err := h.Source.Db().First(&record, "msg_id = ?", in.MsgId).Error; err != nil {
		return
	}

	var clientIds []int64
	if record.DialogType == entity.ChatPrivateMode {
		for _, uid := range [2]int{record.UserId, record.ReceiverId} {
			ids := h.ClientStorage.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(uid))
			clientIds = append(clientIds, ids...)
		}
	} else if record.DialogType == entity.ChatGroupMode {
		clientIds = h.RoomStorage.All(ctx, &cache.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: entity.RoomImGroup,
			Number:   strconv.Itoa(record.ReceiverId),
			Sid:      h.Conf.ServerId(),
		})
	}
	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(entity.PushEventImMessageRevoke, map[string]any{
		"dialog_type": record.DialogType,
		"sender_id":   record.UserId,
		"receiver_id": record.ReceiverId,
		"msg_id":      record.MsgId,
		"text":        "Данное сообщение удалено",
	})
	socket.Session.Chat.Write(c)
}
