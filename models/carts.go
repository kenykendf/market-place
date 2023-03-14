package models

import "gorm.io/gorm"

type Carts struct {
	gorm.Model
	OrderID      uint8
	ItemCode     string `json:"item_code" gorm:"type:varchar(100);column:item_code"`
	ItemName     string `json:"item_name" gorm:"type:varchar(255);column:item_name"`
	Quantity     int    `json:"quantity" gorm:"column:quantity"`
	CategoriesID uint8  `json:"categories_id" gorm:"column:categories_id"`
}
