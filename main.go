package main

import (
	"github.com/Bruce/my-blog/database"
	"github.com/Bruce/my-blog/helpers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("数据库连接初始化失败：%v", err)
	}

	if err := helpers.MigrateData(database.DB); err != nil {
		log.Fatalf("数据库表创建失败：%v", err)
	}

	// 启动gin
	r := gin.Default()

	// 配置路由
	helpers.InitRoute(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("gin启动失败：%v", err)
	}

}
