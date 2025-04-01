package chat

import (
	"context"
	"log"
	"voo.su/internal/constant"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/socket"
)

type MessageRead struct {
	Event   string `json:"event"`
	Content struct {
		Ids        []int `json:"ids"`
		ReceiverId int   `json:"receiver_id"`
	} `json:"content"`
}

func (h *Handler) onReadMessage(ctx context.Context, client socket.IClient, data []byte) {
	var in MessageRead
	if err := jsonutil.Decode(data, &in); err != nil {
		log.Printf("onReadMessage json decode err: %s", err)
		return
	}

	items := make([]model.MessageRead, 0, len(in.Content.Ids))
	for _, msgId := range in.Content.Ids {
		items = append(items, model.MessageRead{
			MessageId:  msgId,
			UserId:     client.Uid(),
			ReceiverId: in.Content.ReceiverId,
		})
	}

	if err := h.MemberUseCase.Source.Postgres().Create(items).Error; err != nil {
		log.Fatalf("onReadMessage Create err: %s", err)
		return
	}

	if err := h.MemberUseCase.Source.Postgres().
		Model(&model.Message{}).
		Where("id in ? AND receiver_id = ? AND is_read = ?", in.Content.Ids, client.Uid(), 0).
		Update("is_read", 1).Error; err != nil {
		log.Printf("onReadMessage Update err:  %s", err)
		return
	}

	h.Redis.Publish(ctx, constant.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessageRead,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   client.Uid(),
			"receiver_id": in.Content.ReceiverId,
			"ids":         in.Content.Ids,
		}),
	}))
}
