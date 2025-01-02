// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"voo.su/internal/constant"
	"voo.su/internal/domain/entity"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	"voo.su/pkg/jsonutil"
	"voo.su/pkg/locale"
)

type ContactRequestUseCase struct {
	Locale locale.ILocale
	Source *infrastructure.Source
}

func NewContactRequestUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
) *ContactRequestUseCase {
	return &ContactRequestUseCase{
		Locale: locale,
		Source: source,
	}
}

type ContactApplyCreateOpt struct {
	UserId   int
	FriendId int
}

func (c *ContactRequestUseCase) Create(ctx context.Context, opt *ContactApplyCreateOpt) error {
	apply := &postgresModel.ContactRequest{
		UserId:   opt.UserId,
		FriendId: opt.FriendId,
	}
	if err := c.Source.Postgres().WithContext(ctx).Create(apply).Error; err != nil {
		return err
	}

	body := map[string]any{
		"event": constant.SubEventContactRequest,
		"data": jsonutil.Encode(map[string]any{
			"apply_id": apply.Id,
			"type":     1,
		}),
	}
	_, err := c.Source.Redis().Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Incr(ctx, fmt.Sprintf("im:contact:apply:%d", opt.FriendId))
		pipe.Publish(ctx, constant.ImTopicChat, jsonutil.Encode(body))
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

type ContactApplyAcceptOpt struct {
	UserId  int
	ApplyId int
}

func (c *ContactRequestUseCase) Accept(ctx context.Context, opt *ContactApplyAcceptOpt) (*postgresModel.ContactRequest, error) {
	db := c.Source.Postgres().WithContext(ctx)
	var applyInfo postgresModel.ContactRequest
	if err := db.First(&applyInfo, "id = ? AND friend_id = ?", opt.ApplyId, opt.UserId).Error; err != nil {
		return nil, err
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		addFriendFunc := func(uid, fid int) error {
			var contact postgresModel.Contact
			err := tx.Where("user_id = ? AND friend_id = ?", uid, fid).First(&contact).Error
			if err == nil {
				return tx.Model(&postgresModel.Contact{}).
					Where("id = ?", contact.Id).
					Updates(&postgresModel.Contact{
						Remark: "",
						Status: 1,
					}).Error
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			return tx.Create(&postgresModel.Contact{
				UserId:   uid,
				FriendId: fid,
				Remark:   "",
				Status:   1,
			}).Error
		}
		var user postgresModel.User
		if err := tx.Select("id", "username").First(&user, applyInfo.FriendId).Error; err != nil {
			return err
		}

		if err := addFriendFunc(applyInfo.UserId, applyInfo.FriendId); err != nil {
			return err
		}

		if err := addFriendFunc(applyInfo.FriendId, applyInfo.UserId); err != nil {
			return err
		}
		return tx.Delete(&postgresModel.ContactRequest{}, "user_id = ? AND friend_id = ?", applyInfo.UserId, applyInfo.FriendId).Error
	})
	return &applyInfo, err
}

type ContactApplyDeclineOpt struct {
	UserId  int
	ApplyId int
}

func (c *ContactRequestUseCase) Decline(ctx context.Context, opt *ContactApplyDeclineOpt) error {
	if err := c.Source.Postgres().
		WithContext(ctx).
		Delete(&postgresModel.ContactRequest{}, "id = ? AND friend_id = ?", opt.ApplyId, opt.UserId).
		Error; err != nil {
		return err
	}

	body := map[string]any{
		"event": constant.SubEventContactRequest,
		"data": jsonutil.Encode(map[string]any{
			"apply_id": opt.ApplyId,
			"type":     2,
		}),
	}
	c.Source.Redis().Publish(ctx, constant.ImTopicChat, jsonutil.Encode(body))

	return nil
}

func (c *ContactRequestUseCase) List(ctx context.Context, uid int) ([]*entity.ApplyItem, error) {
	fields := []string{
		"contact_requests.id",
		"u.username",
		"u.avatar",
		"u.name",
		"u.surname",
		"contact_requests.user_id",
		"contact_requests.friend_id",
		"contact_requests.created_at",
	}
	tx := c.Source.Postgres().WithContext(ctx).
		Table("contact_requests").
		Joins("LEFT JOIN users AS u ON u.id = contact_requests.user_id").
		Where("contact_requests.friend_id = ?", uid).
		Order("contact_requests.id DESC")
	var items []*entity.ApplyItem
	if err := tx.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (c *ContactRequestUseCase) GetApplyUnreadNum(ctx context.Context, uid int) int {
	num, err := c.Source.Redis().Get(ctx, fmt.Sprintf("im:contact:apply:%d", uid)).Int()
	if err != nil {
		return 0
	}

	return num
}

func (c *ContactRequestUseCase) ClearApplyUnreadNum(ctx context.Context, uid int) {
	c.Source.Redis().Del(ctx, fmt.Sprintf("im:contact:apply:%d", uid))
}
