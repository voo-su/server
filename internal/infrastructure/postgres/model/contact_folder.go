package model

import "time"

type ContactFolder struct {
	Id        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL"`
	Name      string    `gorm:"column:name;NOT NULL"`
	Num       int       `gorm:"column:num;default:0;NOT NULL"`
	Sort      int       `gorm:"column:sort;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (ContactFolder) TableName() string {
	return "contact_folders"
}
