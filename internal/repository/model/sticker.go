package model

import "time"

type Sticker struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name;NOT NULL" json:"name"`
	Icon      string    `gorm:"column:icon;NOT NULL" json:"icon"`
	Status    int       `gorm:"column:status;default:0;NOT NULL" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (Sticker) TableName() string {
	return "stickers"
}
