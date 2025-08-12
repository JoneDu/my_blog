package models

// Comment TO
type CommentTO struct {
	Content string `binding:"required"`
	PostID  uint   `binding:"required"`
}
