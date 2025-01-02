// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package usecase

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"voo.su/internal/infrastructure"
	postgresModel "voo.su/internal/infrastructure/postgres/model"
	postgresRepo "voo.su/internal/infrastructure/postgres/repository"
	"voo.su/pkg/locale"
)

type ContactFolderUseCase struct {
	Locale            locale.ILocale
	Source            *infrastructure.Source
	ContactRepo       *postgresRepo.ContactRepository
	ContactFolderRepo *postgresRepo.ContactFolderRepository
}

func NewContactFolderUseCase(
	locale locale.ILocale,
	source *infrastructure.Source,
	contactRepo *postgresRepo.ContactRepository,
	contactFolderRepo *postgresRepo.ContactFolderRepository,
) *ContactFolderUseCase {
	return &ContactFolderUseCase{
		Locale:            locale,
		Source:            source,
		ContactRepo:       contactRepo,
		ContactFolderRepo: contactFolderRepo,
	}
}

func (c *ContactFolderUseCase) Delete(ctx context.Context, id int, uid int) error {
	return c.ContactFolderRepo.Txx(ctx, func(tx *gorm.DB) error {
		res := tx.Delete(&postgresModel.ContactFolder{}, "id = ? AND user_id = ?", id, uid)
		if err := res.Error; err != nil {
			return err
		}

		if res.RowsAffected == 0 {
			return errors.New(c.Locale.Localize("data_not_found"))
		}

		return tx.Table("contacts").
			Where("user_id = ? AND group_id = ?", uid, id).
			UpdateColumn("group_id", 0).
			Error
	})
}

func (c *ContactFolderUseCase) GetUserGroup(ctx context.Context, uid int) ([]*postgresModel.ContactFolder, error) {
	var items []*postgresModel.ContactFolder
	if err := c.Source.Postgres().WithContext(ctx).
		Table("contact_folders").
		Where("user_id = ?", uid).
		Order("sort ASC").
		Scan(&items).
		Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (c *ContactFolderUseCase) MoveGroup(ctx context.Context, uid int, friendId int, groupId int) error {
	contact, err := c.ContactRepo.FindByWhere(ctx, "user_id = ? AND friend_id  = ?", uid, friendId)
	if err != nil {
		return err
	}

	return c.Source.Postgres().WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if contact.FolderId > 0 {
			if err := tx.Table("contact_folders").
				Where("id = ? AND user_id = ?", contact.FolderId, uid).
				Updates(map[string]any{
					"num": gorm.Expr("num - 1"),
				}).Error; err != nil {
				return err
			}
		}

		if err := tx.Table("contacts").
			Where("user_id = ? AND friend_id = ? AND group_id = ?", uid, friendId, contact.FolderId).
			UpdateColumn("group_id", groupId).
			Error; err != nil {
			return err
		}

		return tx.Table("contact_folders").
			Where("id = ? AND user_id = ?", groupId, uid).
			Updates(map[string]any{
				"num": gorm.Expr("num + 1"),
			}).
			Error
	})
}
