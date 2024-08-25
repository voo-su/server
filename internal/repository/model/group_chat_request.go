package model

import "time"

const (
	GroupChatRequestStatusWait   = 1
	GroupChatRequestStatusPass   = 2
	GroupChatRequestStatusRefuse = 3
)

type GroupChatRequest struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	GroupId   int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	Status    int       `gorm:"column:status;default:1;NOT NULL" json:"status"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (GroupChatRequest) TableName() string {
	return "group_chat_requests"
}

type GroupApplyList struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	GroupId   int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	Username  string    `gorm:"column:username;NOT NULL" json:"username"`
	Avatar    string    `gorm:"column:avatar;NOT NULL" json:"avatar"`
}
