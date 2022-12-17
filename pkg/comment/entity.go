package comment

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID              string         `json:"id" db:"id"`
	ReaderId        string         `json:"reader_id" db:"reader_id"`
	PostId          string         `json:"post_id" db:"post_id"`
	ParentCommentId sql.NullString `json:"parent_comment_id,omitempty" db:"parent_id"`
	Content         string         `json:"content" db:"content"`
	PublishedAt     time.Time      `json:"published_at" db:"published_at"`
}

type CommentFromPersistence struct {
	ID              string         `json:"id" db:"id"`
	ReaderId        string         `json:"reader_id" db:"reader_id"`
	ReaderFirstName string         `json:"reader_first_name" db:"reader_first_name"`
	ReaderLastName  string         `json:"reader_last_name" db:"reader_last_name"`
	PostId          string         `json:"post_id" db:"post_id"`
	ParentCommentId sql.NullString `json:"parent_comment_id,omitempty" db:"parent_id"`
	Content         string         `json:"content" db:"content"`
	PublishedAt     time.Time      `json:"published_at" db:"published_at"`
}

type ReaderInfoInsideComment struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type GetAllFromAPostServiceResponse struct {
	ID              string                  `json:"id"`
	PostId          string                  `json:"post_id"`
	ParentCommentId string                  `json:"parent_comment_id,omitempty"`
	Content         string                  `json:"content"`
	PublishedAt     time.Time               `json:"published_at"`
	Reader          ReaderInfoInsideComment `json:"reader"`
}

type CreateCommentFromEndpoint struct {
	ReaderId        string `json:"reader_id" binding:"required,len=36"`
	PostId          string `json:"post_id" binding:"required,len=36"`
	ParentCommentId string `json:"parent_comment_id" binding:"required,len=36"`
	Content         string `json:"content" binding:"required"`
	PublishedAt     string `json:"published_at" binding:"required"`
}

type CreateCommentServiceRequest struct {
	ReaderId        string
	PostId          string
	ParentCommentId string
	Content         string
	PublishedAt     string
}

type UpdateCommentContentDTO struct {
	ID         string `json:"id" db:"id" binding:"required,len=36"`
	NewContent string `json:"new_content" db:"content" binding:"required"`
}
