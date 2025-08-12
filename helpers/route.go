package helpers

import (
	"github.com/Bruce/my-blog/controllers"
	"github.com/Bruce/my-blog/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoute(g *gin.Engine) {
	// 文章评论路由
	comments := g.Group("/comments")
	{
		comments.GET("/post/:id", controllers.GetCommentsByPostId)
		comments.POST("", middlewares.AuthLogin(), controllers.CreateComment)
	}

	// 用户
	users := g.Group("/users")
	{
		users.POST("", controllers.Register)
		users.POST("/login", controllers.Login)
	}

	posts := g.Group("/posts")
	{
		posts.GET("", controllers.GetPosts)
		posts.GET("/:id", controllers.GetPost)
		authPosts := posts.Group("").Use(middlewares.AuthLogin())
		{
			authPosts.POST("", controllers.CreatePost)
		}
	}
}
