package controllers

import (
	"github.com/Bruce/my-blog/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  "请输入正确的用户名，密码，和Email",
			Code: http.StatusBadRequest,
		})
		return
	}
	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}
	user.Password = string(hashedPassword)

	if err := user.CreateUser(); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusOK,
			Data: user.ID,
		})
	}
}
