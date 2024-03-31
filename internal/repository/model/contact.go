package model

import "time"

const (
	ContactStatusNormal = 1
	ContactStatusDelete = 0
)

type Contact struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	FriendId  int       `gorm:"column:friend_id;default:0;NOT NULL" json:"friend_id"`
	Remark    string    `gorm:"column:remark;NOT NULL" json:"remark"`
	Status    int       `gorm:"column:status;default:0;NOT NULL" json:"status"`
	GroupId   int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`
}

func (Contact) TableName() string {
	return "contacts"
}

type ContactListItem struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Avatar   string `grom:"column:avatar" json:"avatar" `
	Name     string `grom:"column:name" json:"name" `
	Surname  string `grom:"column:surname" json:"surname" `
	Gender   uint8  `gorm:"column:gender" json:"gender"`
	About    string `gorm:"column:about" json:"about"`
	Remark   string `gorm:"column:remark" json:"friend_remark"`
	IsOnline int    `json:"isOnline"`
	GroupId  int    `gorm:"column:group_id" json:"group_id"`
}
