package models

import "github.com/jinzhu/gorm"

// Article is model
type Article struct {
	gorm.Model
	Title  string
	Slug   string `gorm:"unique_index"`
	Desc   string `sql:"type:text;"`
	UserID uint
}
