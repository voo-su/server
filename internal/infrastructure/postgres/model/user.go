package model

import "time"

type User struct {
	Id        int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Email     string    `gorm:"column:email;NOT NULL" json:"email"`
	Username  string    `gorm:"column:username;NOT NULL" json:"username"`
	Name      string    `gorm:"column:name;NOT NULL" json:"name"`
	Surname   string    `gorm:"column:surname;NOT NULL" json:"surname"`
	Avatar    string    `gorm:"column:avatar;NOT NULL" json:"avatar"`
	Gender    int       `gorm:"column:gender;default:0;NOT NULL" json:"gender"`
	About     string    `gorm:"column:about;NOT NULL" json:"about"`
	Birthday  string    `gorm:"column:birthday;NOT NULL" json:"birthday"`
	IsBot     int       `gorm:"column:is_bot;default:0;NOT NULL" json:"is_bot"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL" json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
