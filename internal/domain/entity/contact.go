package entity

import "time"

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
	FolderId int    `gorm:"column:folder_id" json:"folder_id"`
}

type ApplyItem struct {
	Id        int       `gorm:"column:id" json:"id"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	FriendId  int       `gorm:"column:friend_id" json:"friend_id"`
	Username  string    `gorm:"column:username" json:"username"`
	Avatar    string    `gorm:"column:avatar" json:"avatar"`
	Name      string    `gorm:"column:name" json:"name"`
	Surname   string    `gorm:"column:surname" json:"surname"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}
