package models

import (
	"github.com/Bruce/my-blog/database"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	PostID  uint
	Post    Post
	UserID  uint
	User    User
}

func GetCommentsByPostId(postId uint) (comments []Comment, err error) {
	// 先判断postid 有没有对应的post
	var post Post
	if err = database.DB.Preload("comments").Take(&post, postId).Error; err != nil {
		return
	}
	// 然后查询出对应的comments
	comments = post.Comments
	return
}
