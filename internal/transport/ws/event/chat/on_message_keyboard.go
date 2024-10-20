package chat

import (
	"context"
	"encoding/json"
	"log"
	"voo.su/internal/entity"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/jsonutil"
)

type KeyboardMessage struct {
	Event   string `json:"event"`
	Content struct {
		SenderID   int `json:"sender_id"`
		ReceiverID int `json:"receiver_id"`
	} `json:"content"`
}

func (h *Handler) onKeyboardMessage(ctx context.Context, _ socket.IClient, data []byte) {
	var in KeyboardMessage
	if err := json.Unmarshal(data, &in); err != nil {
		log.Println("Ошибка в чате при обработке сообщения с клавиатурой: ", err)
		return
	}
	h.Redis.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(map[string]any{
		"event": entity.SubEventImMessageKeyboard,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   in.Content.SenderID,
			"receiver_id": in.Content.ReceiverID,
		}),
	}))
}
