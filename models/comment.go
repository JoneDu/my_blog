package models

import (
	"errors"
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
	var post Post
	if err = database.DB.Preload("comments").Take(&post, postId).Error; err != nil {
		return
	}
	comments = post.Comments
	return
}

// 对文章发评论
func (c *Comment) CreateComment() error {
	userID := c.UserID
	postID := c.PostID
	content := c.Content

	if userID == 0 || postID == 0 || content == "" {
		return errors.New("create comment fail,param error")
	}
	// 根据用户ID，和文章ID查询 文章，是否存在

	return database.DB.Create(c).Error
}
