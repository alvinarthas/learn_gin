package models

import "github.com/jinzhu/gorm"

// User is model
type User struct {
	gorm.Model
	Articles []Article
	UserName string
	FullName string
	Email    string `gorm:"unique_index"`
	SocialID string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
}
