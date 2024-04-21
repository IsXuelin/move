package model

import "github.com/jinzhu/gorm"

type OrderLog struct {
	gorm.Model
	OrderId  string `form:"order_id"`
	Operator string `gorm:"type:VARCHAR(100);not null"`
	DriverId string `gorm:"type:VARCHAR(100);not null"`
}

func (OrderLog) TableName() string {
	return "order_log"
}
