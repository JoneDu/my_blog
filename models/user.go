package models

import (
	"github.com/Bruce/my-blog/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"not null" from:"password" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null" from:"email" json:"email" binding:"required,email"`
	Posts    []Post `form:"-" json:"omitempty"`
}

func (u *User) CreateUser() error {
	return database.DB.Create(u).Error
}
