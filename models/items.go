package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode     string `json:"item_code" gorm:"type:varchar(100);column:item_code"`
	ItemName     string `json:"item_name" gorm:"type:varchar(255);column:item_name"`
	Description  string `json:"description" gorm:"type:varchar(100);column:description"`
	Quantity     int    `json:"quantity" gorm:"column:quantity"`
	CategoriesID uint8  `json:"categories_id" gorm:"column:categories_id"`
}
