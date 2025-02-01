package model

import "time"

type Contact struct {
	Id        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL"`
	FriendId  int       `gorm:"column:friend_id;default:0;NOT NULL"`
	Status    int       `gorm:"column:status;default:0;NOT NULL"`
	FolderId  int       `gorm:"column:group_id;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (Contact) TableName() string {
	return "contacts"
}
