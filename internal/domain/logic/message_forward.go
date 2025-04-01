package logic

import (
	"context"
	"errors"
	"strings"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/strutil"
)

type MessageForward struct {
	Locale       locale.ILocale
	Source       *infrastructure.Source
	SequenceRepo *postgresRepo.SequenceRepository
}

func NewMessageForward(
	locale locale.ILocale,
	source *infrastructure.Source,
	sequence *postgresRepo.SequenceRepository) *MessageForward {
	return &MessageForward{
		Locale:       locale,
		Source:       source,
		SequenceRepo: sequence,
	}
}

type ForwardRecord struct {
	MessageId  int
	ReceiverId int
	ChatType   int
}

func (m *MessageForward) Verify(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) error {
	query := m.Source.Postgres().WithContext(ctx).
		Model(&postgresModel.Message{}).
		Where("id in ?", req.MessageIds)
	if req.Receiver.ChatType == constant.ChatPrivateMode {
		subWhere := m.Source.Postgres().Where("user_id = ? AND receiver_id = ?", uid, req.Receiver.ReceiverId)
		subWhere.Or("user_id = ? AND receiver_id = ?", req.Receiver.ReceiverId, uid)
		query.Where(subWhere)
	}

	query.Where("chat_type = ?", req.Receiver.ChatType).
		Where("msg_type in ?", []int{1, 2, 3, 4, 5, 6, 7, 8, constant.ChatMsgTypeForwarded}).
		Where("is_revoke = ?", 0)

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	if int(count) != len(req.MessageIds) {
		return errors.New("ошибка пересылки сообщения")
	}

	return nil
}

func (m *MessageForward) MultiMergeForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) ([]*ForwardRecord, error) {
	receives := make([]map[string]int, 0)
	for _, userId := range req.Uids {
		receives = append(receives, map[string]int{
			"receiver_id": int(userId),
			"chat_type":   1,
		})
	}
	for _, gid := range req.Gids {
		receives = append(receives, map[string]int{
			"receiver_id": int(gid),
			"chat_type":   2,
		})
	}

	tmpMessages, err := m.aggregation(ctx, req)
	if err != nil {
		return nil, err
	}

	ids := make([]int, 0)
	for _, id := range req.MessageIds {
		ids = append(ids, int(id))
	}

	extra := jsonutil.Encode(entity.MessageExtraForward{
		Ids:      ids,
		Messages: tmpMessages,
	})

	messages := make([]*postgresModel.Message, 0, len(receives))
	for _, item := range receives {
		data := &postgresModel.Message{
			ChatType:   item["chat_type"],
			MsgType:    constant.ChatMsgTypeForwarded,
			UserId:     uid,
			ReceiverId: item["receiver_id"],
			Extra:      extra,
		}
		if data.ChatType == constant.ChatGroupMode {
			data.Sequence = m.SequenceRepo.Get(ctx, 0, data.ReceiverId)
		} else {
			data.Sequence = m.SequenceRepo.Get(ctx, uid, data.ReceiverId)
		}
		messages = append(messages, data)
	}
	if err := m.Source.Postgres().WithContext(ctx).Create(messages).Error; err != nil {
		return nil, err
	}

	list := make([]*ForwardRecord, 0, len(messages))
	for _, record := range messages {
		list = append(list, &ForwardRecord{
			MessageId:  record.Id,
			ReceiverId: record.ReceiverId,
			ChatType:   record.ChatType,
		})
	}

	return list, nil
}

func (m *MessageForward) MultiSplitForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) ([]*ForwardRecord, error) {
	var (
		receives = make([]map[string]int, 0)
		messages = make([]*postgresModel.Message, 0)
		db       = m.Source.Postgres().WithContext(ctx)
	)
	for _, userId := range req.Uids {
		receives = append(receives, map[string]int{
			"receiver_id": int(userId),
			"chat_type":   constant.ChatTypePrivate})
	}

	for _, gid := range req.Gids {
		receives = append(receives, map[string]int{
			"receiver_id": int(gid),
			"chat_type":   constant.ChatTypeGroup,
		})
	}

	if err := db.Model(&postgresModel.Message{}).Where("id IN ?", req.MessageIds).Scan(&messages).Error; err != nil {
		return nil, err
	}

	items := make([]*postgresModel.Message, 0, len(receives)*len(messages))
	messagesLen := int64(len(messages))
	for _, v := range receives {
		var sequences []int64
		if v["chat_type"] == constant.ChatTypeGroup {
			sequences = m.SequenceRepo.BatchGet(ctx, 0, v["receiver_id"], messagesLen)
		} else {
			sequences = m.SequenceRepo.BatchGet(ctx, uid, v["receiver_id"], messagesLen)
		}

		for i, item := range messages {
			items = append(items, &postgresModel.Message{
				ChatType:   v["chat_type"],
				MsgType:    item.MsgType,
				UserId:     uid,
				ReceiverId: v["receiver_id"],
				Content:    item.Content,
				Sequence:   sequences[i],
				Extra:      item.Extra,
			})
		}
	}
	if err := db.Create(items).Error; err != nil {
		return nil, err
	}

	list := make([]*ForwardRecord, 0, len(items))
	for _, item := range items {
		list = append(list, &ForwardRecord{
			MessageId:  item.Id,
			ReceiverId: item.ReceiverId,
			ChatType:   item.ChatType,
		})
	}

	return list, nil
}

type forwardItem struct {
	MsgType  int    `json:"msg_type"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

func (m *MessageForward) aggregation(ctx context.Context, req *v1Pb.ForwardMessageRequest) ([]map[string]any, error) {
	ids := req.MessageIds
	if len(ids) > 3 {
		ids = ids[:3]
	}

	query := m.Source.Postgres().WithContext(ctx).
		Table("messages").
		Joins("LEFT JOIN users ON users.id = messages.user_id").
		Where("messages.id IN ?", ids)
	rows := make([]*forwardItem, 0, 3)
	if err := query.Limit(3).Scan(&rows).Error; err != nil {
		return nil, err
	}

	data := make([]map[string]any, 0)
	for _, row := range rows {
		item := map[string]any{
			"username": row.Username,
		}
		switch row.MsgType {
		case constant.ChatMsgTypeText:
			item["text"] = strutil.MtSubstr(strings.TrimSpace(row.Content), 0, 30)
		case constant.ChatMsgTypeCode:
			item["text"] = m.Locale.Localize("message_with_code")
		case constant.ChatMsgTypeImage:
			item["text"] = m.Locale.Localize("photo")
		case constant.ChatMsgTypeAudio:
			item["text"] = m.Locale.Localize("audio_recording")
		case constant.ChatMsgTypeVideo:
			item["text"] = m.Locale.Localize("video")
		case constant.ChatMsgTypeFile:
			item["text"] = m.Locale.Localize("file")
		case constant.ChatMsgTypeLocation:
			item["text"] = m.Locale.Localize("location_message")
		}
		data = append(data, item)
	}

	return data, nil
}
