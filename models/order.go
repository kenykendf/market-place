package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID      uint `json:"user_id" gorm:"column:user_id"`
	OrderStatus bool `json:"order_status" gorm:"type:boolean"`
	Carts       *[]Carts
}
