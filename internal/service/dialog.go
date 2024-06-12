package service

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
	"voo.su/internal/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type DialogService struct {
	*repo.Source
	Dialog              *repo.Dialog
	GroupChatMemberRepo *repo.GroupChatMember
}

func NewDialogService(
	source *repo.Source,
	dialog *repo.Dialog,
	groupChatMemberRepo *repo.GroupChatMember,
) *DialogService {
	return &DialogService{
		source,
		dialog,
		groupChatMemberRepo,
	}
}

type RemoveRecordListOpt struct {
	UserId     int
	DialogType int
	ReceiverId int
	RecordIds  string
}

func (d *DialogService) DeleteRecordList(ctx context.Context, opt *RemoveRecordListOpt) error {
	var (
		db      = d.Source.Db().WithContext(ctx)
		findIds []int64
		ids     = sliceutil.Unique(sliceutil.ParseIds(opt.RecordIds))
	)
	if opt.DialogType == entity.ChatPrivateMode {
		subQuery := db.Where("user_id = ? and receiver_id = ?", opt.UserId, opt.ReceiverId).
			Or("user_id = ? and receiver_id = ?", opt.ReceiverId, opt.UserId)
		db.Model(&model.Message{}).
			Where("id in ?", ids).
			Where("dialog_type = ?", entity.ChatPrivateMode).
			Where(subQuery).
			Pluck("id", &findIds)
	} else {
		if !d.GroupChatMemberRepo.IsMember(ctx, opt.ReceiverId, opt.UserId, false) {
			return entity.ErrPermissionDenied
		}

		db.Model(&model.Message{}).
			Where("id in ? and dialog_type = ?", ids, entity.ChatGroupMode).
			Pluck("id", &findIds)
	}
	if len(ids) != len(findIds) {
		return errors.New("ошибка удаления")
	}

	items := make([]*model.MessageDelete, 0, len(ids))
	for _, val := range ids {
		items = append(items, &model.MessageDelete{
			RecordId:  val,
			UserId:    opt.UserId,
			CreatedAt: time.Now(),
		})
	}

	return db.Create(items).Error
}

func (d *DialogService) List(ctx context.Context, uid int) ([]*model.SearchDialogSession, error) {
	fields := []string{
		"d.id", "d.dialog_type", "d.receiver_id", "d.updated_at", "d.is_disturb",
		"d.is_top", "d.is_bot", "u.avatar as user_avatar", "u.username",
		"g.group_name", "g.avatar as group_avatar", "u.name", "u.surname",
	}

	query := d.Source.Db().WithContext(ctx).Table("dialogs d")
	query.Joins("LEFT JOIN users AS u ON d.receiver_id = u.id AND d.dialog_type = 1")
	query.Joins("LEFT JOIN group_chats AS g ON d.receiver_id = g.id AND d.dialog_type = 2")
	query.Where("d.user_id = ? and d.is_delete = 0", uid)
	query.Order("d.updated_at DESC")

	var items []*model.SearchDialogSession
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

type DialogCreateOpt struct {
	UserId     int
	DialogType int
	ReceiverId int
	IsBoot     bool
}

func (d *DialogService) Create(ctx context.Context, opt *DialogCreateOpt) (*model.Dialog, error) {
	result, err := d.Dialog.FindByWhere(ctx, "dialog_type = ? and user_id = ? and receiver_id = ?", opt.DialogType, opt.UserId, opt.ReceiverId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		result = &model.Dialog{
			DialogType: opt.DialogType,
			UserId:     opt.UserId,
			ReceiverId: opt.ReceiverId,
		}
		if opt.IsBoot {
			result.IsBot = 1
		}
		d.Source.Db().WithContext(ctx).Create(result)
	} else {
		result.IsTop = 0
		result.IsDelete = 0
		result.IsDisturb = 0
		if opt.IsBoot {
			result.IsBot = 1
		}
		d.Source.Db().WithContext(ctx).Save(result)
	}

	return result, nil
}

func (d *DialogService) Delete(ctx context.Context, uid int, id int) error {
	_, err := d.Dialog.UpdateWhere(ctx, map[string]any{"is_delete": 1, "updated_at": time.Now()}, "id = ? and user_id = ?", id, uid)
	return err
}

type DialogSessionTopOpt struct {
	UserId int
	Id     int
	Type   int
}

func (d *DialogService) Top(ctx context.Context, opt *DialogSessionTopOpt) error {
	_, err := d.Dialog.UpdateWhere(ctx, map[string]any{
		"is_top":     strutil.BoolToInt(opt.Type == 1),
		"updated_at": time.Now(),
	}, "id = ? and user_id = ?", opt.Id, opt.UserId)
	return err
}

type DialogSessionDisturbOpt struct {
	UserId     int
	DialogType int
	ReceiverId int
	IsDisturb  int
}

func (d *DialogService) Disturb(ctx context.Context, opt *DialogSessionDisturbOpt) error {
	_, err := d.Dialog.UpdateWhere(ctx, map[string]any{
		"is_disturb": opt.IsDisturb,
		"updated_at": time.Now(),
	}, "user_id = ? and receiver_id = ? and dialog_type = ?", opt.UserId, opt.ReceiverId, opt.DialogType)
	return err
}

func (d *DialogService) BatchAddList(ctx context.Context, uid int, values map[string]int) {
	ctime := timeutil.DateTime()
	data := make([]string, 0)
	for k, v := range values {
		if v == 0 {
			continue
		}
		value := strings.Split(k, "_")
		if len(value) != 2 {
			continue
		}
		data = append(data, fmt.Sprintf("(%s, %d, %s, '%s', '%s')", value[0], uid, value[1], ctime, ctime))
	}
	if len(data) == 0 {
		return
	}

	d.Source.Db().WithContext(ctx).Exec(fmt.Sprintf("INSERT INTO dialogs (dialog_type, user_id, receiver_id, created_at, updated_at) VALUES %s ON DUPLICATE KEY UPDATE is_delete = 0, updated_at = '%s'", strings.Join(data, ","), ctime))
}
