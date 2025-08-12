package database

import (
	"github.com/Bruce/my-blog/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func InitDB() error {
	config := conf.LoadConfig()

	db, err := gorm.Open(mysql.Open(config.DbDsn), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatal(err)
		return err
	}
	DB = db
	return nil
}
