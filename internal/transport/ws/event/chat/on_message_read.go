package chat

import (
	"context"
	"log"
	"voo.su/internal/entity"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/jsonutil"
)

type DialogReadMessage struct {
	Event   string `json:"event"`
	Content struct {
		MsgIds     []int `json:"msg_id"`
		ReceiverId int   `json:"receiver_id"`
	} `json:"content"`
}

func (h *Handler) onReadMessage(ctx context.Context, client socket.IClient, data []byte) {
	var in DialogReadMessage
	if err := jsonutil.Decode(data, &in); err != nil {
		log.Println("Чат onReadMessage ошибка: ", err)
		return
	}

	h.MemberService.Db().Model(&model.Message{}).
		Where("id in ? and receiver_id = ? and is_read = 0", in.Content.MsgIds, client.Uid()).
		Update("is_read", 1)
	h.Redis.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventImMessageRead,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   client.Uid(),
			"receiver_id": in.Content.ReceiverId,
			"ids":         in.Content.MsgIds,
		}),
	}))
}
