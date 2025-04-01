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
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	redisRepo "voo.su/internal/infrastructure/redis/repository"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
	"voo.su/pkg/sliceutil"
	"voo.su/pkg/strutil"
	"voo.su/pkg/timeutil"
)

type ChatUseCase struct {
	Locale              locale.ILocale
	Source              *infrastructure.Source
	ChatRepo            *postgresRepo.ChatRepository
	GroupChatMemberRepo *postgresRepo.GroupChatMemberRepository
	RedisLockRepo       *redisRepo.RedisLockCacheRepository
	ClientCacheRepo     *redisRepo.ClientCacheRepository
	MessageCacheRepo    *redisRepo.MessageCacheRepository
	UnreadCacheRepo     *redisRepo.UnreadCacheRepository
}

func NewChatUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	chatRepo *postgresRepo.ChatRepository,
	groupChatMemberRepo *postgresRepo.GroupChatMemberRepository,
	redisLockRepo *redisRepo.RedisLockCacheRepository,
	clientCacheRepo *redisRepo.ClientCacheRepository,
	messageCacheRepo *redisRepo.MessageCacheRepository,
	unreadCacheRepository *redisRepo.UnreadCacheRepository,
) *ChatUseCase {
	return &ChatUseCase{
		Locale:              locale,
		Source:              source,
		ChatRepo:            chatRepo,
		GroupChatMemberRepo: groupChatMemberRepo,
		RedisLockRepo:       redisLockRepo,
		ClientCacheRepo:     clientCacheRepo,
		MessageCacheRepo:    messageCacheRepo,
		UnreadCacheRepo:     unreadCacheRepository,
	}
}

func (c *ChatUseCase) List(ctx context.Context, uid int) ([]*entity.SearchChat, error) {
	fields := []string{
		"c.id",
		"c.chat_type",
		"c.receiver_id",
		"c.updated_at",
		"c.is_disturb",
		"c.notify_mute_until",
		"c.notify_show_previews",
		"c.notify_silent",
		"c.is_top",
		"c.is_bot",
		"u.avatar as user_avatar",
		"u.username",
		"g.group_name",
		"g.avatar as group_avatar",
		"u.name",
		"u.surname",
	}

	query := c.Source.Postgres().WithContext(ctx).
		Select(fields).
		Table("chats c").
		Joins("LEFT JOIN users AS u ON c.receiver_id = u.id AND c.chat_type = 1").
		Joins("LEFT JOIN group_chats AS g ON c.receiver_id = g.id AND c.chat_type = 2").
		Where("c.user_id = ? AND c.is_delete = 0", uid).
		Order("c.updated_at DESC")

	var items []*entity.SearchChat
	if err := query.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

type CreateChatOpt struct {
	UserId     int
	ChatType   int
	ReceiverId int
	IsBoot     bool
}

func (c *ChatUseCase) Create(ctx context.Context, opt *CreateChatOpt) (*postgresModel.Chat, error) {
	result, err := c.ChatRepo.FindByWhere(ctx, "chat_type = ? AND user_id = ? AND receiver_id = ?", opt.ChatType, opt.UserId, opt.ReceiverId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		result = &postgresModel.Chat{
			ChatType:   opt.ChatType,
			UserId:     opt.UserId,
			ReceiverId: opt.ReceiverId,
		}
		if opt.IsBoot {
			result.IsBot = 1
		}
		c.Source.Postgres().WithContext(ctx).Create(result)
	} else {
		result.IsTop = 0
		result.IsDelete = 0
		result.IsDisturb = 0
		if opt.IsBoot {
			result.IsBot = 1
		}
		c.Source.Postgres().WithContext(ctx).Save(result)
	}

	return result, nil
}

func (c *ChatUseCase) Delete(ctx context.Context, uid int, id int) error {
	_, err := c.ChatRepo.UpdateWhere(ctx, map[string]any{
		"is_delete":  1,
		"updated_at": time.Now(),
	}, "id = ? AND user_id = ?", id, uid)
	return err
}

type RemoveRecordListOpt struct {
	UserId     int
	ChatType   int
	ReceiverId int
	Ids        []int64
}

func (c *ChatUseCase) DeleteRecordList(ctx context.Context, opt *RemoveRecordListOpt) error {
	var (
		db      = c.Source.Postgres().WithContext(ctx)
		findIds []int64
		ids     = sliceutil.Unique(opt.Ids)
	)
	if opt.ChatType == constant.ChatPrivateMode {
		subQuery := db.Where("user_id = ? AND receiver_id = ?", opt.UserId, opt.ReceiverId).
			Or("user_id = ? AND receiver_id = ?", opt.ReceiverId, opt.UserId)
		db.Model(&postgresModel.Message{}).
			Where("id IN ?", ids).
			Where("chat_type = ?", constant.ChatPrivateMode).
			Where(subQuery).
			Pluck("id", &findIds)
	} else {
		if !c.GroupChatMemberRepo.IsMember(ctx, opt.ReceiverId, opt.UserId, false) {
			return constant.ErrPermissionDenied
		}

		db.Model(&postgresModel.Message{}).
			Where("id in ? AND chat_type = ?", ids, constant.ChatGroupMode).
			Pluck("id", &findIds)
	}
	if len(ids) != len(findIds) {
		return errors.New(c.Locale.Localize("deletion_error"))
	}

	items := make([]*postgresModel.MessageDelete, 0, len(ids))
	for _, val := range ids {
		items = append(items, &postgresModel.MessageDelete{
			MessageId: int(val),
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
	ChatType   int
	ReceiverId int
	IsDisturb  int
}

func (c *ChatUseCase) Disturb(ctx context.Context, opt *ChatDisturbOpt) error {
	_, err := c.ChatRepo.UpdateWhere(ctx, map[string]any{
		"is_disturb": opt.IsDisturb,
		"updated_at": time.Now(),
	}, "user_id = ? AND receiver_id = ? AND chat_type = ?", opt.UserId, opt.ReceiverId, opt.ChatType)
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

	c.Source.Postgres().WithContext(ctx).Exec(fmt.Sprintf("INSERT INTO chats (chat_type, user_id, receiver_id, created_at, updated_at) VALUES %s ON DUPLICATE KEY UPDATE is_delete = 0, updated_at = '%s'", strings.Join(data, ","), ctime))
}

func (c *ChatUseCase) Collect(ctx context.Context, uid int, messageId int64) error {
	var message postgresModel.Message
	if err := c.Source.Postgres().First(&message, messageId).Error; err != nil {
		return err
	}

	if message.MsgType != constant.ChatMsgTypeImage {
		return errors.New(c.Locale.Localize("cannot_favorite_message"))
	}

	if message.IsRevoke == 1 {
		return errors.New(c.Locale.Localize("cannot_favorite_message"))
	}

	if message.ChatType == constant.ChatPrivateMode {
		if message.UserId != uid && message.ReceiverId != uid {
			return constant.ErrPermissionDenied
		}
	} else if message.ChatType == constant.ChatGroupMode {
		if !c.GroupChatMemberRepo.IsMember(ctx, message.ReceiverId, uid, true) {
			return constant.ErrPermissionDenied
		}
	}

	var file entity.MessageExtraImage
	if err := jsonutil.Decode(message.Extra, &file); err != nil {
		return err
	}

	return c.Source.Postgres().Create(&postgresModel.StickerItem{
		UserId:     uid,
		Url:        file.Url,
		FileSuffix: file.Suffix,
		FileSize:   file.Size,
	}).Error
}

func (c *ChatUseCase) GetNotifySettings(ctx context.Context, chatType int, uid int, chatId int64) (*entity.NotifySettings, error) {
	notify, err := c.ChatRepo.FindByWhere(ctx, "user_id = ? AND receiver_id = ? AND chat_type = ?", uid, chatId, chatType)
	if err != nil {
		return nil, err
	}

	return &entity.NotifySettings{
		MuteUntil:    notify.NotifyMuteUntil,
		ShowPreviews: notify.NotifyShowPreviews,
		Silent:       notify.NotifySilent,
	}, nil
}

func (c *ChatUseCase) UpdateNotifySettings(ctx context.Context, chatType int, uid int, chatId int64, data *entity.NotifySettings) error {
	// TODO
	disturb := 0
	if data.MuteUntil > 0 {
		disturb = 1
	}
	_, err := c.ChatRepo.UpdateWhere(ctx, map[string]any{
		"is_disturb":           disturb,
		"notify_mute_until":    data.MuteUntil,
		"notify_show_previews": data.ShowPreviews,
		"notify_silent":        data.Silent,
	}, "user_id = ? AND receiver_id = ? AND chat_type = ?", uid, chatId, chatType)
	return err

}
