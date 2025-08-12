package models

import (
	"errors"
	"github.com/Bruce/my-blog/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" form:"username" json:"username" binding:"required"`
	Password string `gorm:"not null" from:"password" json:"password" binding:"required"`
	Email    string `gorm:"unique;not null" from:"email" json:"email" binding:"required,email"`
	Posts    []Post `form:"-"`
}

func (u *User) CreateUser() error {
	return database.DB.Create(u).Error
}

func GetUserByEmail(email string) (user User, err error) {
	database.DB.Where("email = ?", email).First(&user)
	if user.ID == 0 {
		err = errors.New("user not found")
	}
	return
}
