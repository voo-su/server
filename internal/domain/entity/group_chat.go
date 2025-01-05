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
	Id        int       `gorm:"column:id"`
	GroupId   int       `gorm:"column:group_id"`
	UserId    int       `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_atL"`
	Username  string    `gorm:"column:username"`
	Avatar    string    `gorm:"column:avatar"`
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
