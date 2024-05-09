package model

import "time"

type ContactFolder struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Name      string    `gorm:"column:name;NOT NULL" json:"remark"`
	Num       int       `gorm:"column:num;default:0;NOT NULL" json:"num"`
	Sort      int       `gorm:"column:sort;default:0;NOT NULL" json:"sort"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (ContactFolder) TableName() string {
	return "contact_folders"
}
