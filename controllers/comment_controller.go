package controllers

import (
	"github.com/Bruce/my-blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCommentsByPostId(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
		return
	}

	if comments, err := models.GetCommentsByPostId(uint(postId)); err != nil {
		c.JSON(http.StatusOK, models.Response{
			Msg:  err.Error(),
			Code: http.StatusBadRequest,
		})
	} else {
		c.JSON(http.StatusOK, models.Response{
			Data: comments,
			Code: http.StatusOK,
		})
	}

}
