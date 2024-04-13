package cron

import (
	"context"
	"gorm.io/gorm"
	"path"
	"time"
	"voo.su/internal/repository/model"
	"voo.su/pkg/filesystem"
)

type ClearTmpFile struct {
	DB         *gorm.DB
	Filesystem *filesystem.Filesystem
}

func NewClearTmpFile(db *gorm.DB, fileSystem *filesystem.Filesystem) *ClearTmpFile {
	return &ClearTmpFile{DB: db, Filesystem: fileSystem}
}

func (c *ClearTmpFile) Name() string {
	return "clear.tmp.file"
}

func (c *ClearTmpFile) Spec() string {
	return "20 1 * * *"
}

func (c *ClearTmpFile) Enable() bool {
	return true
}

func (c *ClearTmpFile) Handle(ctx context.Context) error {
	lastId, size := 0, 100
	for {
		items := make([]*model.Split, 0)
		err := c.DB.Model(&model.Split{}).
			Where("id > ? and type = 1 and drive = 1 and created_at <= ?", lastId, time.Now().Add(-24*time.Hour)).
			Order("id asc").
			Limit(size).
			Scan(&items).Error
		if err != nil {
			return err
		}

		for _, item := range items {
			list := make([]*model.Split, 0)
			c.DB.Table("splits").
				Where("user_id = ? and upload_id = ? and type = 2", item.UserId, item.UploadId).
				Scan(&list)
			for _, value := range list {
				if err := c.Filesystem.Local.Delete(value.Path); err == nil {
					c.DB.Delete(model.Split{}, value.Id)
				}
			}

			if len(list) > 0 {
				_ = c.Filesystem.Local.DeleteDir(path.Dir(list[0].Path))
			}

			if err := c.Filesystem.Local.Delete(item.Path); err == nil {
				c.DB.Delete(model.Split{}, item.Id)
			}
		}

		if len(items) == size {
			lastId = items[size-1].Id
		} else {
			break
		}
	}

	return nil
}
