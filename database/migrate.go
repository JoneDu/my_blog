package database

import (
	"github.com/Bruce/my-blog/models"
	"gorm.io/gorm"
)

func MigrateData(db *gorm.DB) error {
	return db.AutoMigrate(&models.Comment{}, &models.Post{}, &models.User{})
}
