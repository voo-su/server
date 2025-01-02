// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package entity

import "time"

type GroupItem struct {
	Id          int    `json:"id"`
	GroupName   string `json:"group_name"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	Leader      int    `json:"leader"`
	IsDisturb   int    `json:"is_disturb"`
	CreatorId   int    `json:"creator_id"`
}

type GroupApplyList struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	GroupId   int       `gorm:"column:group_id;default:0;NOT NULL" json:"group_id"`
	UserId    int       `gorm:"column:user_id;default:0;NOT NULL" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	Username  string    `gorm:"column:username;NOT NULL" json:"username"`
	Avatar    string    `gorm:"column:avatar;NOT NULL" json:"avatar"`
}

type MemberItem struct {
	Id       string `json:"id"`
	UserId   int    `json:"user_id"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Gender   int    `json:"gender"`
	About    string `json:"about"`
	Leader   int    `json:"leader"`
	IsMute   int    `json:"is_mute"`
	//UserCard string `json:"user_card"`
}

type SearchAdsItem struct {
	Id           int       `json:"id" grom:"column:id"`
	CreatorId    int       `json:"creator_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	IsTop        int       `json:"is_top"`
	IsConfirm    int       `json:"is_confirm"`
	ConfirmUsers string    `json:"confirm_users"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Avatar       string    `json:"avatar"`
	Username     string    `json:"username"`
}
