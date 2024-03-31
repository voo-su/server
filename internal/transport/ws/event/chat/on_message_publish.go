package chat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"log"
	"voo.su/api/pb/v1"
	"voo.su/pkg/core/socket"
)

var publishMapping map[string]handle

func (h *Handler) onPublish(ctx context.Context, client socket.IClient, data []byte) {
	if publishMapping == nil {
		publishMapping = make(map[string]handle)
		publishMapping["text"] = h.onTextMessage
		publishMapping["code"] = h.onCodeMessage
		publishMapping["location"] = h.onLocationMessage
		publishMapping["sticker"] = h.onStickerMessage
		publishMapping["vote"] = h.onVoteMessage
		publishMapping["image"] = h.onImageMessage
		publishMapping["file"] = h.onFileMessage
	}
	val, err := sonic.Get(data, "content.type")
	if err == nil {
		return
	}
	typeValue, _ := val.String()
	if call, ok := publishMapping[typeValue]; ok {
		call(ctx, client, data)
	} else {
		log.Printf("Событие чата: onPublish %s неизвестный тип сообщения\n", typeValue)
	}
}

type TextMessage struct {
	AckId   string                    `json:"ack_id"`
	Event   string                    `json:"event"`
	Content api_v1.TextMessageRequest `json:"content"`
}

func (h *Handler) onTextMessage(ctx context.Context, client socket.IClient, data []byte) {
	var in TextMessage
	if err := json.Unmarshal(data, &in); err != nil {
		log.Println("Ошибка в чате при получении текстового сообщения: ", err)
		return
	}
	if in.Content.GetContent() == "" || in.Content.GetReceiver() == nil {
		return
	}
	err := h.message.SendText(ctx, client.Uid(), &api_v1.TextMessageRequest{
		Content: in.Content.Content,
		Receiver: &api_v1.MessageReceiver{
			DialogType: in.Content.Receiver.DialogType,
			ReceiverId: in.Content.Receiver.ReceiverId,
		},
	})
	if err != nil {
		log.Printf("Ошибка в чате при получении текстового сообщения: %s", err.Error())
		return
	}
	if len(in.AckId) == 0 {
		return
	}
	if err = client.Write(&socket.ClientResponse{Sid: in.AckId, Event: "ack"}); err != nil {
		log.Printf("Ошибка подтверждения в чате при получении текстового сообщения: %s", err.Error())
	}
}

type CodeMessage struct {
	AckId   string                    `json:"ack_id"`
	Event   string                    `json:"event"`
	Content api_v1.CodeMessageRequest `json:"content"`
}

func (h *Handler) onCodeMessage(ctx context.Context, client socket.IClient, data []byte) {
	var m CodeMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Println("Ошибка в чате при получении текстового сообщения: ", err)
		return
	}
	if m.Content.GetReceiver() == nil {
		return
	}

	err := h.message.SendCode(ctx, client.Uid(), &api_v1.CodeMessageRequest{
		Lang: m.Content.Lang,
		Code: m.Content.Code,
		Receiver: &api_v1.MessageReceiver{
			DialogType: m.Content.Receiver.DialogType,
			ReceiverId: m.Content.Receiver.ReceiverId,
		},
	})
	if err != nil {
		log.Printf("Ошибка в чате при получении текстового сообщения: %s", err.Error())
		return
	}

	if len(m.AckId) == 0 {
		return
	}
	if err = client.Write(&socket.ClientResponse{Sid: m.AckId, Event: "ack"}); err != nil {
		log.Printf("Ошибка подтверждения в чате при получении текстового сообщения: %s", err.Error())
	}
}

type StickerMessage struct {
	MsgId   string                       `json:"msg_id"`
	Event   string                       `json:"event"`
	Content api_v1.StickerMessageRequest `json:"content"`
}

func (h *Handler) onStickerMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m StickerMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Println("Ошибка в чате при обработке сообщения с смайликом: ", err)
		return
	}
	fmt.Println("[onStickerMessage] Новое сообщение ", string(data))
}

type ImageMessage struct {
	MsgId   string                     `json:"msg_id"`
	Event   string                     `json:"event"`
	Content api_v1.ImageMessageRequest `json:"content"`
}

func (h *Handler) onImageMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m ImageMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Println("Ошибка в чате при обработке сообщения с изображением: ", err)
		return
	}
	fmt.Println("[onImageMessage] Новое сообщение ", string(data))
}

type FileMessage struct {
	MsgId   string                     `json:"msg_id"`
	Event   string                     `json:"event"`
	Content api_v1.ImageMessageRequest `json:"content"`
}

func (h *Handler) onFileMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m FileMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Println("Ошибка в чате при обработке файла: ", err)
		return
	}
	fmt.Println("[onFileMessage] Новое сообщение ", string(data))
}

type LocationMessage struct {
	MsgId   string                        `json:"msg_id"`
	Event   string                        `json:"event"`
	Content api_v1.LocationMessageRequest `json:"content"`
}

func (h *Handler) onLocationMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m LocationMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Println("Ошибка в чате при обработке сообщения с местоположением: ", err)
		return
	}
	fmt.Println("[onLocationMessage] Новое сообщение ", string(data))
}

type VoteMessage struct {
	MsgId   string                    `json:"msg_id"`
	Event   string                    `json:"event"`
	Content api_v1.VoteMessageRequest `json:"content"`
}

func (h *Handler) onVoteMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m VoteMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Println("Ошибка в чате при обработке голосового сообщения: ", err)
		return
	}
	fmt.Println("[onVoteMessage] Новое сообщение ", string(data))
}
