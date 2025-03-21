package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"voo.su/internal/constant"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/socket"
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
		log.Fatalf("onKeyboardMessage json decode err: %s", err.Error())
		return
	}

	_data := jsonutil.Encode(map[string]any{
		"event": constant.SubEventImMessageKeyboard,
		"data": jsonutil.Encode(map[string]any{
			"sender_id":   in.Content.SenderID,
			"receiver_id": in.Content.ReceiverID,
		}),
	})

	if err := h.Nats.Publish(constant.ImTopicChat, []byte(_data)); err != nil {
		fmt.Println(err)
	}

	h.Redis.Publish(ctx, constant.ImTopicChat, _data)
}
