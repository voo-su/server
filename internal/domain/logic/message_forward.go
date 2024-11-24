package logic

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"strings"
	v1Pb "voo.su/api/http/pb/v1"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/strutil"
)

type MessageForwardLogic struct {
	DB       *gorm.DB
	Sequence *repo.Sequence
}

func NewMessageForwardLogic(db *gorm.DB, sequence *repo.Sequence) *MessageForwardLogic {
	return &MessageForwardLogic{
		DB:       db,
		Sequence: sequence,
	}
}

type ForwardRecord struct {
	RecordId   int
	ReceiverId int
	DialogType int
}

func (m *MessageForwardLogic) Verify(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) error {
	query := m.DB.WithContext(ctx).Model(&model.Message{})
	query.Where("id in ?", req.MessageIds)
	if req.Receiver.DialogType == constant.ChatPrivateMode {
		subWhere := m.DB.Where("user_id = ? AND receiver_id = ?", uid, req.Receiver.ReceiverId)
		subWhere.Or("user_id = ? AND receiver_id = ?", req.Receiver.ReceiverId, uid)
		query.Where(subWhere)
	}

	query.Where("dialog_type = ?", req.Receiver.DialogType)
	query.Where("msg_type in ?", []int{1, 2, 3, 4, 5, 6, 7, 8, constant.ChatMsgTypeForward})
	query.Where("is_revoke = ?", 0)

	var count int64
	if err := query.Count(&count).Error; err != nil {
		return err
	}
	if int(count) != len(req.MessageIds) {
		return errors.New("ошибка пересылки сообщения")
	}

	return nil
}

func (m *MessageForwardLogic) MultiMergeForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) ([]*ForwardRecord, error) {
	receives := make([]map[string]int, 0)
	for _, userId := range req.Uids {
		receives = append(receives, map[string]int{"receiver_id": int(userId), "dialog_type": 1})
	}
	for _, gid := range req.Gids {
		receives = append(receives, map[string]int{"receiver_id": int(gid), "dialog_type": 2})
	}

	tmpRecords, err := m.aggregation(ctx, req)
	if err != nil {
		return nil, err
	}

	ids := make([]int, 0)
	for _, id := range req.MessageIds {
		ids = append(ids, int(id))
	}

	extra := jsonutil.Encode(entity.DialogRecordExtraForward{
		MsgIds:  ids,
		Records: tmpRecords,
	})

	records := make([]*model.Message, 0, len(receives))
	for _, item := range receives {
		data := &model.Message{
			MsgId:      strutil.NewMsgId(),
			DialogType: item["dialog_type"],
			MsgType:    constant.ChatMsgTypeForward,
			UserId:     uid,
			ReceiverId: item["receiver_id"],
			Extra:      extra,
		}
		if data.DialogType == constant.ChatGroupMode {
			data.Sequence = m.Sequence.Get(ctx, 0, data.ReceiverId)
		} else {
			data.Sequence = m.Sequence.Get(ctx, uid, data.ReceiverId)
		}
		records = append(records, data)
	}
	if err := m.DB.WithContext(ctx).Create(records).Error; err != nil {
		return nil, err
	}

	list := make([]*ForwardRecord, 0, len(records))
	for _, record := range records {
		list = append(list, &ForwardRecord{
			RecordId:   record.Id,
			ReceiverId: record.ReceiverId,
			DialogType: record.DialogType,
		})
	}

	return list, nil
}

func (m *MessageForwardLogic) MultiSplitForward(ctx context.Context, uid int, req *v1Pb.ForwardMessageRequest) ([]*ForwardRecord, error) {
	var (
		receives = make([]map[string]int, 0)
		records  = make([]*model.Message, 0)
		db       = m.DB.WithContext(ctx)
	)
	for _, userId := range req.Uids {
		receives = append(receives, map[string]int{
			"receiver_id": int(userId),
			"dialog_type": constant.DialogRecordDialogTypePrivate})
	}

	for _, gid := range req.Gids {
		receives = append(receives, map[string]int{
			"receiver_id": int(gid),
			"dialog_type": constant.DialogRecordDialogTypeGroup,
		})
	}

	if err := db.Model(&model.Message{}).Where("id IN ?", req.MessageIds).Scan(&records).Error; err != nil {
		return nil, err
	}

	items := make([]*model.Message, 0, len(receives)*len(records))
	recordsLen := int64(len(records))
	for _, v := range receives {
		var sequences []int64
		if v["dialog_type"] == constant.DialogRecordDialogTypeGroup {
			sequences = m.Sequence.BatchGet(ctx, 0, v["receiver_id"], recordsLen)
		} else {
			sequences = m.Sequence.BatchGet(ctx, uid, v["receiver_id"], recordsLen)
		}

		for i, item := range records {
			items = append(items, &model.Message{
				MsgId:      strutil.NewMsgId(),
				DialogType: v["dialog_type"],
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
			RecordId:   item.Id,
			ReceiverId: item.ReceiverId,
			DialogType: item.DialogType,
		})
	}

	return list, nil
}

type forwardItem struct {
	MsgType  int    `json:"msg_type"`
	Content  string `json:"content"`
	Username string `json:"username"`
}

func (m *MessageForwardLogic) aggregation(ctx context.Context, req *v1Pb.ForwardMessageRequest) ([]map[string]any, error) {
	rows := make([]*forwardItem, 0, 3)
	query := m.DB.WithContext(ctx).Table("messages")
	query.Joins("LEFT JOIN users on users.id = messages.user_id")
	ids := req.MessageIds
	if len(ids) > 3 {
		ids = ids[:3]
	}

	query.Where("messages.id in ?", ids)
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
			item["text"] = "Сообщение с кодом"
		case constant.ChatMsgTypeImage:
			item["text"] = "Фотография"
		case constant.ChatMsgTypeAudio:
			item["text"] = "Аудиозапись"
		case constant.ChatMsgTypeVideo:
			item["text"] = "Видео"
		case constant.ChatMsgTypeFile:
			item["text"] = "Файл"
		case constant.ChatMsgTypeLocation:
			item["text"] = "Сообщение с местоположением"
		}
		data = append(data, item)
	}

	return data, nil
}
