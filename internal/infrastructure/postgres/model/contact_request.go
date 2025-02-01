package model

import "time"

type ContactRequest struct {
	Id        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL"`
	FriendId  int       `gorm:"column:friend_id;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
}

func (ContactRequest) TableName() string {
	return "contact_requests"
}
