package controllers

import (
	"github.com/Bruce/my-blog/database"
	"github.com/Bruce/my-blog/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
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

func GetPostById(c *gin.Context) {

}

func GetPosts(c *gin.Context) {
	type simplePost struct {
		gorm.Model
		Title   string
		Content string
		UserID  uint
	}
	var posts []simplePost
	getPosts, err := models.GetPosts()
	for i := range getPosts {
		posts = append(posts, simplePost{
			Model: gorm.Model{
				ID:        getPosts[i].ID,
				CreatedAt: getPosts[i].CreatedAt,
				UpdatedAt: getPosts[i].UpdatedAt,
			},
			Title:   getPosts[i].Title,
			Content: getPosts[i].Content,
			UserID:  getPosts[i].UserID,
		})
	}
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusOK,
			Data: posts,
		})
	}
}

func GetPost(c *gin.Context) {
	pId := c.Param("id")
	id, err := strconv.Atoi(pId)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	var post models.Post
	err = database.DB.First(&post, id).Error
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusInternalServerError,
			Msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Code: http.StatusOK,
			Data: post,
		})
	}

}

func DeletePostById(c *gin.Context) {

}
