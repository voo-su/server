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

func NewContactService(
	source *repo.Source,
	contact *repo.Contact,
) *ContactService {
	return &ContactService{
		Source:  source,
		contact: contact,
	}
}

func (c *ContactService) UpdateRemark(ctx context.Context, uid int, friendId int, remark string) error {
	_, err := c.contact.UpdateWhere(ctx, map[string]any{
		"remark": remark,
	}, "user_id = ? and friend_id = ?", uid, friendId)
	if err == nil {
		_ = c.contact.SetFriendRemark(ctx, uid, friendId, remark)
	}

	return err
}

func (c *ContactService) Delete(ctx context.Context, uid, friendId int) error {
	return c.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return tx.Table("contacts").
			Where("user_id = ? and friend_id = ?", uid, friendId).
			Update("status", model.ContactStatusDelete).Error
	})
}

func (c *ContactService) List(ctx context.Context, uid int) ([]*model.ContactListItem, error) {
	tx := c.contact.Model(ctx)
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
	c.contact.Model(ctx).Where("user_id = ? and status = ?", uid, model.ContactStatusNormal).Pluck("friend_id", &ids)
	return ids
}
