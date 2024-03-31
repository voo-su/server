package service

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type ContactService struct {
	*repo.Source
	contact *repo.Contact
}

func NewContactService(source *repo.Source, contact *repo.Contact) *ContactService {
	return &ContactService{Source: source, contact: contact}
}

func (s *ContactService) UpdateRemark(ctx context.Context, uid int, friendId int, remark string) error {
	_, err := s.contact.UpdateWhere(ctx, map[string]any{"remark": remark}, "user_id = ? and friend_id = ?", uid, friendId)
	if err == nil {
		_ = s.contact.SetFriendRemark(ctx, uid, friendId, remark)
	}

	return err
}

func (s *ContactService) Delete(ctx context.Context, uid, friendId int) error {
	find, err := s.contact.FindByWhere(ctx, "user_id = ? and friend_id = ?", uid, friendId)
	if err != nil {
		return err
	}

	return s.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if find.GroupId > 0 {
			err := tx.Table("contact_groups").
				Where("id = ? and user_id = ?", find.GroupId, uid).
				Updates(map[string]any{"num": gorm.Expr("num - 1")}).Error

			if err != nil {
				return err
			}
		}
		return tx.Table("contacts").
			Where("user_id = ? and friend_id = ?", uid, friendId).
			Update("status", model.ContactStatusDelete).Error
	})
}

func (s *ContactService) List(ctx context.Context, uid int) ([]*model.ContactListItem, error) {
	tx := s.contact.Model(ctx)
	tx.Select([]string{
		"u.id",
		"u.username",
		"u.avatar",
		"u.name",
		"u.surname",
		"u.about",
		"u.gender",
		"contacts.remark",
		"contacts.group_id",
	})
	tx.Joins("INNER JOIN users AS u ON u.id = contacts.friend_id")
	tx.Where("contacts.user_id = ? AND contacts.status = ?", uid, model.ContactStatusNormal)
	var items []*model.ContactListItem
	if err := tx.Scan(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (s *ContactService) GetContactIds(ctx context.Context, uid int) []int64 {
	var ids []int64
	s.contact.Model(ctx).Where("user_id = ? and status = ?", uid, model.ContactStatusNormal).Pluck("friend_id", &ids)
	return ids
}

func (s *ContactService) MoveGroup(ctx context.Context, uid int, friendId int, groupId int) error {
	contact, err := s.contact.FindByWhere(ctx, "user_id = ? and friend_id  = ?", uid, friendId)
	if err != nil {
		return err
	}

	return s.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if contact.GroupId > 0 {
			err := tx.Table("contact_groups").
				Where("id = ? and user_id = ?", contact.GroupId, uid).
				Updates(map[string]any{
					"num": gorm.Expr("num - 1"),
				}).Error

			if err != nil {
				return err
			}
		}
		err := tx.Table("contacts").
			Where("user_id = ? and friend_id = ? and group_id = ?", uid, friendId, contact.GroupId).
			UpdateColumn("group_id", groupId).
			Error
		if err != nil {
			return err
		}

		return tx.Table("contact_groups").Where("id = ? and user_id = ?", groupId, uid).Updates(map[string]any{
			"num": gorm.Expr("num + 1"),
		}).Error
	})
}
