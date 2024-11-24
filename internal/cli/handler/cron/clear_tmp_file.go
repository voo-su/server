package cron

import (
	"context"
	"gorm.io/gorm"
	"path"
	"time"
	"voo.su/internal/repository/model"
	"voo.su/pkg/minio"
)

type ClearTmpFile struct {
	DB    *gorm.DB
	Minio minio.IMinio
}

func NewClearTmpFile(db *gorm.DB, minio minio.IMinio) *ClearTmpFile {
	return &ClearTmpFile{
		DB:    db,
		Minio: minio,
	}
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
		if err := c.DB.Model(&model.Split{}).
			Where("id > ? AND type = 1 AND drive = 1 AND created_at <= ?", lastId, time.Now().Add(-24*time.Hour)).
			Order("id asc").
			Limit(size).
			Scan(&items).
			Error; err != nil {
			return err
		}

		for _, item := range items {
			list := make([]*model.Split, 0)
			c.DB.Table("splits").
				Where("user_id = ? AND upload_id = ? AND type = 2", item.UserId, item.UploadId).
				Scan(&list)

			for _, value := range list {
				_ = c.Minio.Delete(c.Minio.BucketPublicName(), value.Path)
				c.DB.Delete(model.Split{}, value.Id)
			}

			if len(list) > 0 {
				_ = c.Minio.Delete(c.Minio.BucketPrivateName(), path.Dir(item.Path))
			}

			if err := c.Minio.Delete(c.Minio.BucketPrivateName(), item.Path); err == nil {
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
