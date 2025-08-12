package middlewares

import (
	"github.com/Bruce/my-blog/conf"
	"github.com/Bruce/my-blog/jwt"
	"github.com/Bruce/my-blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, models.Response{
				Code: http.StatusUnauthorized,
				Msg:  "没有正确携带Token",
			})
			return
		}
		tokenString := strings.Split(authHeader, " ")[1]
		if tokenString == "" {
			c.JSON(http.StatusOK, models.Response{
				Code: http.StatusInternalServerError,
				Msg:  "Token 格式有误",
			})
			return
		}

		config := conf.LoadConfig()
		cliam, err := jwt.ParseToken(tokenString, config.JWTSecret)
		if err != nil {
			c.JSON(http.StatusOK, models.Response{
				Code: http.StatusUnauthorized,
				Msg:  "无效令牌",
			})
			return
		}
		c.Set("username", cliam.Username)
		c.Set("userId", cliam.UserId)
		c.Next()
	}
}
