package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/constant"
	"voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/logger"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/socket"
	"voo.su/pkg/timeutil"
)

type ConsumeContactStatus struct {
	Status int `json:"status"`
	UserId int `json:"user_id"`
}

func (h *Handler) onConsumeContactStatus(ctx context.Context, body []byte) {
	var in ConsumeContactStatus
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeContactStatus json decode err: %s", err)
		return
	}
	contactIds := h.ContactUseCase.GetContactIds(ctx, in.UserId)

	clientIds := make([]int64, 0)
	for _, uid := range sliceutil.Unique(contactIds) {
		ids := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.FormatInt(uid, 10))
		if len(ids) > 0 {
			clientIds = append(clientIds, ids...)
		}
	}
	if len(clientIds) == 0 {
		return
	}

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetMessage(constant.PushEventContactStatus, in)
	socket.Session.Chat.Write(c)
}

type ConsumeContactApply struct {
	ApplyId int `json:"apply_id"`
	Type    int `json:"type"`
}

func (h *Handler) onConsumeContactApply(ctx context.Context, body []byte) {
	var in ConsumeContactApply
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeContactApply json decode err: %s", err.Error())
		return
	}

	var apply model.ContactRequest
	if err := h.ContactUseCase.Source.Postgres().First(&apply, in.ApplyId).Error; err != nil {
		return
	}

	clientIds := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(apply.FriendId))
	if len(clientIds) == 0 {
		return
	}

	var user model.User
	if err := h.ContactUseCase.Source.Postgres().First(&user, apply.FriendId).Error; err != nil {
		return
	}

	data := map[string]any{}
	data["sender_id"] = apply.UserId
	data["receiver_id"] = apply.FriendId
	data["friend"] = map[string]any{
		"username":   user.Username,
		"created_at": timeutil.FormatDatetime(apply.CreatedAt),
	}
	c := socket.NewSenderContent()
	c.SetAck(true)
	c.SetReceive(clientIds...)
	c.SetMessage(constant.PushEventContactRequest, data)
	socket.Session.Chat.Write(c)
}
