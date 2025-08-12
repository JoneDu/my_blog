package controllers

import (
	"github.com/Bruce/my-blog/conf"
	"github.com/Bruce/my-blog/jwt"
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
			Code: http.StatusInternalServerError,
		})
		return
	}
	user.Password = string(hashedPassword)

	if err := user.CreateUser(); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  err.Error(),
			Code: http.StatusInternalServerError,
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusCreated,
			Data: user.ID,
			Msg:  "恭喜！！注册成功！",
		})
	}
}

func Login(c *gin.Context) {
	var userLogin models.UserLogin
	if err := c.ShouldBind(&userLogin); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  "输入正确的password，email",
			Code: http.StatusBadRequest,
		})
		return
	}
	// 查询用户，交验密码是否一致
	user, err := models.GetUserByEmail(userLogin.Email)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  err.Error(),
			Code: http.StatusInternalServerError,
		})
		return
	}
	// 如果没有通过提示登录失败
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  "密码错误！登录失败！",
			Code: http.StatusUnauthorized,
		})
		return
	}

	// 成功登录则返回一个token，有效期30分钟
	config := conf.LoadConfig()
	token, err := jwt.GenerateToken(user.Username, user.ID, config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  err.Error(),
			Code: http.StatusInternalServerError,
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusOK,
			Data: token,
		})
	}

}
