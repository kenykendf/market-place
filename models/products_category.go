package models

import "gorm.io/gorm"

type Categories struct {
	gorm.Model
	CategoriesName string `json:"categories_name" gorm:"type:varchar(256);column:categories_name"`
	Item           []Item
}
