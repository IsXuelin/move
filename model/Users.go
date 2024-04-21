package model

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	UserId      string `form:"user_id"`
	UserName    string `gorm:"type:VARCHAR(100);not null"`
	Category    string `gorm:"type:VARCHAR(100);not null"`
	PhoneNumber int64  `gorm:"type:INT"`
}

func (Users) TableName() string {
	return "users"
}
