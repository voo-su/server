package usecase

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type ChatUseCase struct {
	*repo.Source
	ChatRepo            *repo.Chat
	GroupChatMemberRepo *repo.GroupChatMember
}

func NewChatUseCase(
	source *repo.Source,
	chatRepo *repo.Chat,
	groupChatMemberRepo *repo.GroupChatMember,
) *ChatUseCase {
	return &ChatUseCase{
		source,
		chatRepo,
		groupChatMemberRepo,
	}
}

func (c *ChatUseCase) List(ctx context.Context, uid int) ([]*entity.SearchChat, error) {
	fields := []string{
		"c.id",
		"c.dialog_type",
		"c.receiver_id",
		"c.updated_at",
		"c.is_disturb",
		"c.is_top",
		"c.is_bot",
		"u.avatar as user_avatar",
		"u.username",
		"g.group_name",
		"g.avatar as group_avatar",
		"u.name",
		"u.surname",
	}

	query := c.Source.Db().WithContext(ctx).Table("chats c")
	query.Joins("LEFT JOIN users AS u ON c.receiver_id = u.id AND c.dialog_type = 1")
	query.Joins("LEFT JOIN group_chats AS g ON c.receiver_id = g.id AND c.dialog_type = 2")
	query.Where("c.user_id = ? AND c.is_delete = 0", uid)
	query.Order("c.updated_at DESC")

	var items []*entity.SearchChat
	if err := query.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

type CreateChatOpt struct {
	UserId     int
	DialogType int
	ReceiverId int
	IsBoot     bool
}

func (c *ChatUseCase) Create(ctx context.Context, opt *CreateChatOpt) (*model.Chat, error) {
	result, err := c.ChatRepo.FindByWhere(ctx, "dialog_type = ? AND user_id = ? AND receiver_id = ?", opt.DialogType, opt.UserId, opt.ReceiverId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		result = &model.Chat{
			DialogType: opt.DialogType,
			UserId:     opt.UserId,
			ReceiverId: opt.ReceiverId,
		}
		if opt.IsBoot {
			result.IsBot = 1
		}
		c.Source.Db().WithContext(ctx).Create(result)
	} else {
		result.IsTop = 0
		result.IsDelete = 0
		result.IsDisturb = 0
		if opt.IsBoot {
			result.IsBot = 1
		}
		c.Source.Db().WithContext(ctx).Save(result)
	}

	return result, nil
}

func (c *ChatUseCase) Delete(ctx context.Context, uid int, id int) error {
	_, err := c.ChatRepo.UpdateWhere(ctx, map[string]any{"is_delete": 1, "updated_at": time.Now()}, "id = ? AND user_id = ?", id, uid)
	return err
}

type RemoveRecordListOpt struct {
	UserId     int
	DialogType int
	ReceiverId int
	RecordIds  string
}

func (c *ChatUseCase) DeleteRecordList(ctx context.Context, opt *RemoveRecordListOpt) error {
	var (
		db      = c.Source.Db().WithContext(ctx)
		findIds []int64
		ids     = sliceutil.Unique(sliceutil.ParseIds(opt.RecordIds))
	)
	if opt.DialogType == constant.ChatPrivateMode {
		subQuery := db.Where("user_id = ? AND receiver_id = ?", opt.UserId, opt.ReceiverId).
			Or("user_id = ? AND receiver_id = ?", opt.ReceiverId, opt.UserId)
		db.Model(&model.Message{}).
			Where("id in ?", ids).
			Where("dialog_type = ?", constant.ChatPrivateMode).
			Where(subQuery).
			Pluck("id", &findIds)
	} else {
		if !c.GroupChatMemberRepo.IsMember(ctx, opt.ReceiverId, opt.UserId, false) {
			return constant.ErrPermissionDenied
		}

		db.Model(&model.Message{}).
			Where("id in ? AND dialog_type = ?", ids, constant.ChatGroupMode).
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

type ChatTopOpt struct {
	UserId int
	Id     int
	Type   int
}

func (c *ChatUseCase) Top(ctx context.Context, opt *ChatTopOpt) error {
	_, err := c.ChatRepo.UpdateWhere(ctx, map[string]any{
		"is_top":     strutil.BoolToInt(opt.Type == 1),
		"updated_at": time.Now(),
	}, "id = ? AND user_id = ?", opt.Id, opt.UserId)
	return err
}

type ChatDisturbOpt struct {
	UserId     int
	DialogType int
	ReceiverId int
	IsDisturb  int
}

func (c *ChatUseCase) Disturb(ctx context.Context, opt *ChatDisturbOpt) error {
	_, err := c.ChatRepo.UpdateWhere(ctx, map[string]any{
		"is_disturb": opt.IsDisturb,
		"updated_at": time.Now(),
	}, "user_id = ? AND receiver_id = ? AND dialog_type = ?", opt.UserId, opt.ReceiverId, opt.DialogType)
	return err
}

func (c *ChatUseCase) BatchAddList(ctx context.Context, uid int, values map[string]int) {
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

	c.Source.Db().WithContext(ctx).Exec(fmt.Sprintf("INSERT INTO chats (dialog_type, user_id, receiver_id, created_at, updated_at) VALUES %s ON DUPLICATE KEY UPDATE is_delete = 0, updated_at = '%s'", strings.Join(data, ","), ctime))
}

func (c *ChatUseCase) Collect(ctx context.Context, uid int, recordId int) error {
	var record model.Message
	if err := c.Source.Db().First(&record, recordId).Error; err != nil {
		return err
	}

	if record.MsgType != constant.ChatMsgTypeImage {
		return errors.New("это сообщение нельзя добавить в избранное")
	}

	if record.IsRevoke == 1 {
		return errors.New("это сообщение нельзя добавить в избранное")
	}

	if record.DialogType == constant.ChatPrivateMode {
		if record.UserId != uid && record.ReceiverId != uid {
			return constant.ErrPermissionDenied
		}
	} else if record.DialogType == constant.ChatGroupMode {
		if !c.GroupChatMemberRepo.IsMember(ctx, record.ReceiverId, uid, true) {
			return constant.ErrPermissionDenied
		}
	}

	var file entity.DialogRecordExtraImage
	if err := jsonutil.Decode(record.Extra, &file); err != nil {
		return err
	}

	return c.Source.Db().Create(&model.StickerItem{
		UserId:     uid,
		Url:        file.Url,
		FileSuffix: file.Suffix,
		FileSize:   file.Size,
	}).Error
}
