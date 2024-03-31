package model

import "time"

type ContactRequest struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	FriendId  int       `gorm:"column:friend_id;default:0;NOT NULL" json:"friend_id"`
	Remark    string    `gorm:"column:remark;NOT NULL" json:"remark"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
}

func (ContactRequest) TableName() string {
	return "contact_requests"
}

type ApplyItem struct {
	Id        int       `gorm:"column:id" json:"id"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	FriendId  int       `gorm:"column:friend_id" json:"friend_id"`
	Remark    string    `gorm:"column:remark" json:"remark"`
	Username  string    `gorm:"column:username" json:"username"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Name      string    `gorm:"column:name" json:"name"`
	Surname   string    `gorm:"column:surname" json:"surname"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
