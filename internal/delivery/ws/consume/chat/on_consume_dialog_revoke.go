package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/constant"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	redisModel "voo.su/internal/infrastructure/redis/model"
	"voo.su/pkg/logger"
	"voo.su/pkg/socket"
)

type ConsumeMessageRevoke struct {
	MsgId string `json:"msg_id"`
}

func (h *Handler) onConsumeMessageRevoke(ctx context.Context, body []byte) {
	var in ConsumeMessageRevoke
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeMessageRevoke json decode err: %s", err.Error())
		return
	}

	var record postgresModel.Message
	if err := h.Source.Postgres().First(&record, "msg_id = ?", in.MsgId).Error; err != nil {
		return
	}

	var clientIds []int64
	if record.ChatType == constant.ChatPrivateMode {
		for _, uid := range [2]int{record.UserId, record.ReceiverId} {
			ids := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(uid))
			clientIds = append(clientIds, ids...)
		}
	} else if record.ChatType == constant.ChatGroupMode {
		clientIds = h.RoomCache.All(ctx, &redisModel.RoomOption{
			Channel:  socket.Session.Chat.Name(),
			RoomType: constant.RoomImGroup,
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
	c.SetMessage(constant.PushEventImMessageRevoke, map[string]any{
		"chat_type":   record.ChatType,
		"sender_id":   record.UserId,
		"receiver_id": record.ReceiverId,
		"msg_id":      record.MsgId,
		"text":        h.Locale.Localize("message_deleted"),
	})
	socket.Session.Chat.Write(c)
}
