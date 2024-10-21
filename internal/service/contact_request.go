package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"voo.su/internal/entity"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
	"voo.su/pkg/jsonutil"
)

type ContactRequestService struct {
	*repo.Source
}

func NewContactApplyService(source *repo.Source) *ContactRequestService {
	return &ContactRequestService{Source: source}
}

type ContactApplyCreateOpt struct {
	UserId   int
	FriendId int
}

func (s *ContactRequestService) Create(ctx context.Context, opt *ContactApplyCreateOpt) error {
	apply := &model.ContactRequest{
		UserId:   opt.UserId,
		FriendId: opt.FriendId,
	}
	if err := s.Source.Db().WithContext(ctx).Create(apply).Error; err != nil {
		return err
	}

	body := map[string]any{
		"event": entity.SubEventContactRequest,
		"data": jsonutil.Encode(map[string]any{
			"apply_id": apply.Id,
			"type":     1,
		}),
	}
	_, _ = s.Source.Redis().Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Incr(ctx, fmt.Sprintf("im:contact:apply:%d", opt.FriendId))
		pipe.Publish(ctx, entity.ImTopicChat, jsonutil.Encode(body))
		return nil
	})
	return nil
}

type ContactApplyAcceptOpt struct {
	UserId  int
	ApplyId int
}

func (s *ContactRequestService) Accept(ctx context.Context, opt *ContactApplyAcceptOpt) (*model.ContactRequest, error) {
	db := s.Source.Db().WithContext(ctx)
	var applyInfo model.ContactRequest
	if err := db.First(&applyInfo, "id = ? and friend_id = ?", opt.ApplyId, opt.UserId).Error; err != nil {
		return nil, err
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		addFriendFunc := func(uid, fid int) error {
			var contact model.Contact
			err := tx.Where("user_id = ? and friend_id = ?", uid, fid).First(&contact).Error
			if err == nil {
				return tx.
					Model(&model.Contact{}).
					Where("id = ?", contact.Id).
					Updates(&model.Contact{
						Remark: "",
						Status: 1,
					}).Error
			}
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				return err
			}

			return tx.Create(&model.Contact{
				UserId:   uid,
				FriendId: fid,
				Remark:   "",
				Status:   1,
			}).Error
		}
		var user model.User
		if err := tx.Select("id", "username").First(&user, applyInfo.FriendId).Error; err != nil {
			return err
		}

		if err := addFriendFunc(applyInfo.UserId, applyInfo.FriendId); err != nil {
			return err
		}

		if err := addFriendFunc(applyInfo.FriendId, applyInfo.UserId); err != nil {
			return err
		}
		return tx.Delete(&model.ContactRequest{}, "user_id = ? and friend_id = ?", applyInfo.UserId, applyInfo.FriendId).Error
	})
	return &applyInfo, err
}

type ContactApplyDeclineOpt struct {
	UserId  int
	ApplyId int
}

func (s *ContactRequestService) Decline(ctx context.Context, opt *ContactApplyDeclineOpt) error {
	err := s.Source.Db().WithContext(ctx).Delete(&model.ContactRequest{}, "id = ? and friend_id = ?", opt.ApplyId, opt.UserId).Error
	if err != nil {
		return err
	}

	body := map[string]any{
		"event": entity.SubEventContactRequest,
		"data": jsonutil.Encode(map[string]any{
			"apply_id": opt.ApplyId,
			"type":     2,
		}),
	}
	s.Source.Redis().Publish(ctx, entity.ImTopicChat, jsonutil.Encode(body))

	return nil
}

func (s *ContactRequestService) List(ctx context.Context, uid int) ([]*model.ApplyItem, error) {
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
	tx := s.Source.Db().WithContext(ctx).Table("contact_requests")
	tx.Joins("LEFT JOIN users AS u ON u.id = contact_requests.user_id")
	tx.Where("contact_requests.friend_id = ?", uid)
	tx.Order("contact_requests.id DESC")
	var items []*model.ApplyItem
	if err := tx.Select(fields).Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (s *ContactRequestService) GetApplyUnreadNum(ctx context.Context, uid int) int {
	num, err := s.Source.Redis().Get(ctx, fmt.Sprintf("im:contact:apply:%d", uid)).Int()
	if err != nil {
		return 0
	}

	return num
}

func (s *ContactRequestService) ClearApplyUnreadNum(ctx context.Context, uid int) {
	s.Source.Redis().Del(ctx, fmt.Sprintf("im:contact:apply:%d", uid))
}
