package model

import "time"

type GroupChatRequest struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	GroupId   int       `gorm:"column:group_id;default:0;NOT NULL"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL"`
	Status    int       `gorm:"column:status;default:1;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (GroupChatRequest) TableName() string {
	return "group_chat_requests"
}
