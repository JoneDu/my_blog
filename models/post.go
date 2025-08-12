package models

import (
	"github.com/Bruce/my-blog/database"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Content  string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	User     User
	Comments []Comment
}

func (p *Post) CreatePost() error {
	return database.DB.Create(p).Error
}

func GetPosts() ([]Post, error) {
	var posts []Post
	err := database.DB.Omit("deleted_at").Find(&posts).Error
	return posts, err
}
