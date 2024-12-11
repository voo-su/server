package model

import "time"

type StickerItem struct {
	Id          int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	StickerId   int       `gorm:"column:sticker_id;default:0;NOT NULL" json:"sticker_id"`
	UserId      int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Description string    `gorm:"column:description;NOT NULL" json:"description"`
	Url         string    `gorm:"column:url;NOT NULL" json:"url"`
	FileSuffix  string    `gorm:"column:file_suffix;NOT NULL" json:"file_suffix"`
	FileSize    int       `gorm:"column:file_size;default:0;NOT NULL" json:"file_size"`
	CreatedAt   time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (StickerItem) TableName() string {
	return "sticker_items"
}
