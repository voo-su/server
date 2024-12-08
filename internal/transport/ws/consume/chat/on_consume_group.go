package chat

import (
	"context"
	"encoding/json"
	"strconv"
	"voo.su/internal/constant"
	"voo.su/internal/repository/cache"
	model2 "voo.su/internal/repository/model"
	"voo.su/pkg/logger"
	"voo.su/pkg/socket"
)

type ConsumeGroupJoin struct {
	Gid  int   `json:"group_id"`
	Type int   `json:"type"`
	Uids []int `json:"uids"`
}

type ConsumeGroupApply struct {
	GroupId int `json:"group_id"`
	UserId  int `json:"user_id"`
}

func (h *Handler) onConsumeGroupJoin(ctx context.Context, body []byte) {
	var in ConsumeGroupJoin
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("onConsumeGroupJoin Ошибка при декодировании: ", err.Error())
		return
	}

	sid := h.Conf.ServerId()
	for _, uid := range in.Uids {

		ids := h.ClientCache.GetUidFromClientIds(ctx, sid, socket.Session.Chat.Name(), strconv.Itoa(uid))
		for _, cid := range ids {
			opt := &cache.RoomOption{
				Channel:  socket.Session.Chat.Name(),
				RoomType: constant.RoomImGroup,
				Number:   strconv.Itoa(in.Gid),
				Sid:      h.Conf.ServerId(),
				Cid:      cid,
			}

			if in.Type == 2 {
				_ = h.RoomCache.Del(ctx, opt)
			} else {
				_ = h.RoomCache.Add(ctx, opt)
			}
		}
	}
}

func (h *Handler) onConsumeGroupApply(ctx context.Context, body []byte) {
	var in ConsumeGroupApply
	if err := json.Unmarshal(body, &in); err != nil {
		logger.Errorf("[ChatSubscribe] onConsumeGroupApply Unmarshal ошибка: ", err.Error())
		return
	}

	var groupMember model2.GroupChatMember
	if err := h.Source.Db().First(&groupMember, "group_id = ? AND leader = ?", in.GroupId, 2).Error; err != nil {
		return
	}

	var groupDetail model2.GroupChat
	if err := h.Source.Db().First(&groupDetail, in.GroupId).Error; err != nil {
		return
	}

	var user model2.User
	if err := h.Source.Db().First(&user, in.UserId).Error; err != nil {
		return
	}

	data := make(map[string]any)
	data["group_name"] = groupDetail.Name
	data["username"] = user.Username

	clientIds := h.ClientCache.GetUidFromClientIds(ctx, h.Conf.ServerId(), socket.Session.Chat.Name(), strconv.Itoa(groupMember.UserId))

	c := socket.NewSenderContent()
	c.SetReceive(clientIds...)
	c.SetMessage(constant.PushEventGroupChatRequest, data)

	socket.Session.Chat.Write(c)
}
