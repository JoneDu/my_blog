package controllers

import (
	"github.com/Bruce/my-blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostTO struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func CreatePost(c *gin.Context) {
	var postTo PostTO
	if err := c.BindJSON(&postTo); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}

	var post models.Post
	post.Title = postTo.Title
	post.Content = postTo.Content
	userId, _ := c.Get("userId")
	post.UserID = userId.(uint)

	// 创建Post
	if err := post.CreatePost(); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusOK,
			Data: post.ID,
		})
	}
}
