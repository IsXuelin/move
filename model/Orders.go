package model

import (
	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	OrderId        string `form:"order_id"`
	StartLatitude  string `gorm:"type:VARCHAR(100);not null"`
	StartLongitude string `gorm:"type:VARCHAR(100);not null"`
	EndLatitude    string `gorm:"type:VARCHAR(100);not null"`
	EndLongitude   string `gorm:"type:VARCHAR(100);not null"`
	Distance       int    `gorm:"type:INT;not null"`
	Status         int    `gorm:"type:INT;not null"`
	DriverId       string `gorm:"type:VARCHAR(100)"`
	CustomerID     string `gorm:"type:VARCHAR(100)"`
}

func (Orders) TableName() string {
	return "orders"
}
