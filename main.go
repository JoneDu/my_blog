package main

import (
	"github.com/Bruce/my-blog/database"
	"log"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库连接初始化失败：%v", err)
	}

	if err := database.MigrateData(database.DB); err != nil {
		log.Fatalf("数据库表创建失败：%v", err)
	}

}
