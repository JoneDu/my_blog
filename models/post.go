package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User
	Comments []Comment
}
