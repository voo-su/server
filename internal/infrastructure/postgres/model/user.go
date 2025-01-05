package model

import "time"

type User struct {
	Id        int       `gorm:"primaryKey"`
	Email     string    `gorm:"column:email;NOT NULL"`
	Username  string    `gorm:"column:username;NOT NULL"`
	Name      string    `gorm:"column:name;NOT NULL"`
	Surname   string    `gorm:"column:surname;NOT NULL"`
	Avatar    string    `gorm:"column:avatar;NOT NULL"`
	Gender    int       `gorm:"column:gender;default:0;NOT NULL"`
	About     string    `gorm:"column:about;NOT NULL"`
	Birthday  string    `gorm:"column:birthday;NOT NULL"`
	IsBot     int       `gorm:"column:is_bot;default:0;NOT NULL"`
	CreatedAt time.Time `gorm:"column:created_at;NOT NULL"`
	UpdatedAt time.Time `gorm:"column:updated_at;NOT NULL"`
}

func (User) TableName() string {
	return "users"
}
