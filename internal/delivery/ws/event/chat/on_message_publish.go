package chat

import (
	"context"
	"encoding/json"
	"github.com/bytedance/sonic"
	"log"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/domain/entity"
	"voo.su/pkg/socket"
)

var publishMapping map[string]handle

func (h *Handler) onPublish(ctx context.Context, client socket.IClient, data []byte) {
	if publishMapping == nil {
		publishMapping = make(map[string]handle)
		publishMapping["text"] = h.onTextMessage
		publishMapping["image"] = h.onImageMessage
		publishMapping["vote"] = h.onVoteMessage
		publishMapping["file"] = h.onFileMessage
		publishMapping["sticker"] = h.onStickerMessage
		publishMapping["code"] = h.onCodeMessage
		publishMapping["location"] = h.onLocationMessage
	}
	val, err := sonic.Get(data, "content.type")
	if err == nil {
		return
	}
	typeValue, _ := val.String()
	if call, ok := publishMapping[typeValue]; ok {
		call(ctx, client, data)
	} else {
		log.Printf("onPublish [%s] unknown message type", typeValue)
	}
}

type TextMessage struct {
	AckId   string                    `json:"ack_id"`
	Event   string                    `json:"event"`
	Content entity.TextMessageRequest `json:"content"`
}

func (h *Handler) onTextMessage(ctx context.Context, client socket.IClient, data []byte) {
	var in TextMessage
	if err := json.Unmarshal(data, &in); err != nil {
		log.Printf("onTextMessage json decode err: %s", err)
		return
	}
	if in.Content.Content == "" || in.Content.Receiver == nil {
		return
	}

	if err := h.MessageUseCase.SendText(ctx, client.Uid(), &entity.SendText{
		Content: in.Content.Content,
		Receiver: entity.MessageReceiver{
			ChatType:   in.Content.Receiver.ChatType,
			ReceiverId: in.Content.Receiver.ReceiverId,
		},
	}); err != nil {
		log.Printf("onTextMessage SendText: %s", err.Error())
		return
	}

	if len(in.AckId) == 0 {
		return
	}
	if err := client.Write(&socket.ClientResponse{
		Sid:   in.AckId,
		Event: "ack",
	}); err != nil {
		log.Printf("onTextMessage: %s", err.Error())
	}
}

type StickerMessage struct {
	MsgId   string                     `json:"msg_id"`
	Event   string                     `json:"event"`
	Content v1Pb.StickerMessageRequest `json:"content"`
}

func (h *Handler) onStickerMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m StickerMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Printf("onStickerMessage json decode err: %s", err)
		return
	}
}

type ImageMessage struct {
	MsgId   string                   `json:"msg_id"`
	Event   string                   `json:"event"`
	Content v1Pb.ImageMessageRequest `json:"content"`
}

func (h *Handler) onImageMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m ImageMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Printf("onImageMessage json decode err: %s", err)
		return
	}
}

type FileMessage struct {
	MsgId   string                   `json:"msg_id"`
	Event   string                   `json:"event"`
	Content v1Pb.ImageMessageRequest `json:"content"`
}

func (h *Handler) onFileMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m FileMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Printf("onFileMessage json decode err: %s", err)
		return
	}
}

type VoteMessage struct {
	MsgId   string                  `json:"msg_id"`
	Event   string                  `json:"event"`
	Content v1Pb.VoteMessageRequest `json:"content"`
}

func (h *Handler) onVoteMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m VoteMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Printf("onVoteMessage json decode err: %s", err)
		return
	}
}

type CodeMessage struct {
	AckId   string                  `json:"ack_id"`
	Event   string                  `json:"event"`
	Content v1Pb.CodeMessageRequest `json:"content"`
}

func (h *Handler) onCodeMessage(ctx context.Context, client socket.IClient, data []byte) {
	var m CodeMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Printf("onCodeMessage json decode err: %s", err)
		return
	}
	if m.Content.GetReceiver() == nil {
		return
	}

	if err := h.MessageUseCase.SendCode(ctx, client.Uid(), &v1Pb.CodeMessageRequest{
		Lang: m.Content.Lang,
		Code: m.Content.Code,
		Receiver: &v1Pb.MessageReceiver{
			ChatType:   m.Content.Receiver.ChatType,
			ReceiverId: m.Content.Receiver.ReceiverId,
		},
	}); err != nil {
		log.Printf("onCodeMessage SendCode: %s", err.Error())
		return
	}

	if len(m.AckId) == 0 {
		return
	}
	if err := client.Write(&socket.ClientResponse{
		Sid:   m.AckId,
		Event: "ack",
	}); err != nil {
		log.Printf("onCodeMessage: %s", err.Error())
	}
}

type LocationMessage struct {
	MsgId   string                      `json:"msg_id"`
	Event   string                      `json:"event"`
	Content v1Pb.LocationMessageRequest `json:"content"`
}

func (h *Handler) onLocationMessage(_ context.Context, _ socket.IClient, data []byte) {
	var m LocationMessage
	if err := json.Unmarshal(data, &m); err != nil {
		log.Printf("onLocationMessage json decode err: %s", err)
		return
	}
}
