package model

import "time"

type Sticker struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"column:name;NOT NULL"`
	Icon      string    `gorm:"column:icon;NOT NULL"`
	Status    int       `gorm:"column:status;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (Sticker) TableName() string {
	return "stickers"
}
