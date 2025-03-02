package model

import "time"

type Contact struct {
	Id        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"column:user_id;DEFAULT:0;NOT NULL"`
	FriendId  int       `gorm:"column:friend_id;DEFAULT:0;NOT NULL"`
	Status    int       `gorm:"column:status;DEFAULT:0;NOT NULL"`
	FolderId  int       `gorm:"column:group_id;DEFAULT:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;DEFAULT:CURRENT_TIMESTAMP;NOT NULL"`
}

func (Contact) TableName() string {
	return "contacts"
}
