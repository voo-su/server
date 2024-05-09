package service

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"voo.su/internal/repository/model"
	"voo.su/internal/repository/repo"
)

type ContactFolderService struct {
	*repo.Source
	ContactRepo       *repo.Contact
	ContactFolderRepo *repo.ContactFolder
}

func NewContactFolderService(
	source *repo.Source,
	contactRepo *repo.Contact,
	contactFolderRepo *repo.ContactFolder,
) *ContactFolderService {
	return &ContactFolderService{
		Source:            source,
		ContactRepo:       contactRepo,
		ContactFolderRepo: contactFolderRepo,
	}
}

func (c *ContactFolderService) Delete(ctx context.Context, id int, uid int) error {
	return c.ContactFolderRepo.Txx(ctx, func(tx *gorm.DB) error {
		res := tx.Delete(&model.ContactFolder{}, "id = ? and user_id = ?", id, uid)
		if err := res.Error; err != nil {
			return err
		}
		if res.RowsAffected == 0 {
			return errors.New("данные не существуют")
		}
		return tx.Table("contacts").
			Where("user_id = ? and group_id = ?", uid, id).
			UpdateColumn("group_id", 0).Error
	})
}

func (c *ContactFolderService) GetUserGroup(ctx context.Context, uid int) ([]*model.ContactFolder, error) {
	var items []*model.ContactFolder
	err := c.Source.Db().WithContext(ctx).
		Table("contact_folders").
		Where("user_id = ?", uid).
		Order("sort asc").
		Scan(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (c *ContactFolderService) MoveGroup(ctx context.Context, uid int, friendId int, groupId int) error {
	contact, err := c.ContactRepo.FindByWhere(ctx, "user_id = ? and friend_id  = ?", uid, friendId)
	if err != nil {
		return err
	}

	return c.Source.Db().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if contact.FolderId > 0 {
			err := tx.Table("contact_folders").
				Where("id = ? and user_id = ?", contact.FolderId, uid).
				Updates(map[string]any{
					"num": gorm.Expr("num - 1"),
				}).Error

			if err != nil {
				return err
			}
		}
		err := tx.Table("contacts").
			Where("user_id = ? and friend_id = ? and group_id = ?", uid, friendId, contact.FolderId).
			UpdateColumn("group_id", groupId).
			Error
		if err != nil {
			return err
		}

		return tx.Table("contact_folders").Where("id = ? and user_id = ?", groupId, uid).Updates(map[string]any{
			"num": gorm.Expr("num + 1"),
		}).Error
	})
}
