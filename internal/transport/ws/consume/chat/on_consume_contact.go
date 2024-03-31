package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/entity"
	"voo.su/internal/repository/model"
	"voo.su/pkg/core/socket"
	"voo.su/pkg/logger"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/timeutil"
)

type ConsumeContactStatus struct {
	Status int `json:"status"`
	UserId int `json:"user_id"`
}

func (h *Handler) onConsumeContactStatus(ctx context.Context, body []byte) {
	var in ConsumeContactStatus
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeContactStatus Ошибка при декодировании: ", err.Error())
		return
	}
	contactIds := h.ContactService.GetContactIds(ctx, in.UserId)

	clientIds := make([]int64, 0)
	for _, uid := range sliceutil.Unique(contactIds) {
		ids := h.ClientStorage.GetUidFromClientIds(ctx, h.Config.ServerId(), socket.Session.Chat.Name(), strconv.FormatInt(uid, 10))
		if len(ids) > 0 {
			clientIds = append(clientIds, ids...)
		}
	}
	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetMessage(entity.PushEventContactStatus, in)
	socket.Session.Chat.Write(c)
}

type ConsumeContactApply struct {
	ApplyId int `json:"apply_id"`
	Type    int `json:"type"`
}

func (h *Handler) onConsumeContactApply(ctx context.Context, body []byte) {
	var in ConsumeContactApply
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeContactApply Ошибка при декодировании: ", err.Error())
		return
	}

	var apply model.ContactRequest
	if err := h.ContactService.Db().First(&apply, in.ApplyId).Error; err != nil {
		return
	}

	clientIds := h.ClientStorage.GetUidFromClientIds(ctx, h.Config.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(apply.FriendId))
	if len(clientIds) == 0 {
		return
	}

	var user model.User
	if err := h.ContactService.Db().First(&user, apply.FriendId).Error; err != nil {
		return
	}

	data := map[string]any{}
	data["sender_id"] = apply.UserId
	data["receiver_id"] = apply.FriendId
	data["remark"] = apply.Remark
	data["friend"] = map[string]any{
		"username":   user.Username,
		"remark":     apply.Remark,
		"created_at": timeutil.FormatDatetime(apply.CreatedAt),
	}
	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(entity.PushEventContactRequest, data)
	socket.Session.Chat.Write(c)
}
