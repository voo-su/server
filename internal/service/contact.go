package service

import (
	"context"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type ContactService struct {
	*repo.Source
	Contact *repo.Contact
}

func NewContactService(
	source *repo.Source,
	contact *repo.Contact,
) *ContactService {
	return &ContactService{
		Source:  source,
		Contact: contact,
	}
}

func (c *ContactService) List(ctx context.Context, uid int) ([]*model.ContactListItem, error) {
	tx := c.Contact.Model(ctx)
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

func (c *ContactService) GetContactIds(ctx context.Context, uid int) []int64 {
	var ids []int64
	c.Contact.Model(ctx).
		Where("user_id = ? and status = ?", uid, model.ContactStatusNormal).
		Pluck("friend_id", &ids)
	return ids
}

func (c *ContactService) Delete(ctx context.Context, uid, friendId int) error {
	find, err := c.Contact.FindByWhere(ctx, "user_id = ? and friend_id = ?", uid, friendId)
	if err != nil {
		return err
	}

	return c.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if find.FolderId > 0 {
			err := tx.Table("contact_folders").
				Where("id = ? and user_id = ?", find.FolderId, uid).
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

func (c *ContactService) UpdateRemark(ctx context.Context, uid int, friendId int, remark string) error {
	_, err := c.Contact.UpdateWhere(ctx, map[string]any{
		"remark": remark,
	}, "user_id = ? and friend_id = ?", uid, friendId)
	if err == nil {
		_ = c.Contact.SetFriendRemark(ctx, uid, friendId, remark)
	}

	return err
}
