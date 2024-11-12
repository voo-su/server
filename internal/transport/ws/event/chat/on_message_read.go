package chat

import (
	"context"
	"log"
	"voo.su/internal/constant"
	"voo.su/internal/repository/model"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/logger"
	"voo.su/pkg/socket"
)

type DialogReadMessage struct {
	Event   string `json:"event"`
	Content struct {
		MsgIds     []string `json:"msg_ids"`
		ReceiverId int      `json:"receiver_id"`
	} `json:"content"`
}

func (h *Handler) onReadMessage(ctx context.Context, client socket.IClient, data []byte) {
	var in DialogReadMessage
	if err := jsonutil.Decode(data, &in); err != nil {
		log.Println("Чат onReadMessage ошибка: ", err)
		return
	}

	items := make([]model.MessageRead, 0, len(in.Content.MsgIds))
	for _, msgId := range in.Content.MsgIds {
		items = append(items, model.MessageRead{
			MsgId:      msgId,
			UserId:     client.Uid(),
			ReceiverId: in.Content.ReceiverId,
		})
	}

	if err := h.MemberUseCase.Db().Create(items).Error; err != nil {
		logger.Error("Не удалось выполнить пакетное создание MessageRead", err.Error())
		return
	}

	if err := h.MemberUseCase.Db().
		Model(&model.Message{}).
		Where("msg_id in ? AND receiver_id = ? AND is_read = 0", in.Content.MsgIds, client.Uid()).
		Update("is_read", 1).Error; err != nil {
		log.Println("Чат onReadMessage ошибка: ", err)
		return
	}

	h.Redis.Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessageRead,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   client.Uid(),
			"receiver_id": in.Content.ReceiverId,
			"msg_ids":     in.Content.MsgIds,
		}),
	}))
}
