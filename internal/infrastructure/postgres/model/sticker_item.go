package model

import "time"

type StickerItem struct {
	Id          int       `gorm:"primaryKey"`
	StickerId   int       `gorm:"column:sticker_id;default:0;NOT NULL"`
	UserId      int       `gorm:"column:user_id;default:0;NOT NULL"`
	Description string    `gorm:"column:description;NOT NULL"`
	Url         string    `gorm:"column:url;NOT NULL"`
	FileSuffix  string    `gorm:"column:file_suffix;NOT NULL"`
	FileSize    int       `gorm:"column:file_size;default:0;NOT NULL"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt   time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (StickerItem) TableName() string {
	return "sticker_items"
}
