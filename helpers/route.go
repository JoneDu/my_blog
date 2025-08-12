package helpers

import (
	"github.com/Bruce/my-blog/controllers"
	"github.com/gin-gonic/gin"
)

func InitRoute(g *gin.Engine) {
	// 文章评论路由
	comments := g.Group("/comments")
	{
		comments.GET("/post/:id", controllers.GetCommentsByPostId)
		comments.POST("", controllers.CreateComment)
	}

	// 用户
	users := g.Group("/users")
	{
		users.POST("", controllers.Register)
	}
}
