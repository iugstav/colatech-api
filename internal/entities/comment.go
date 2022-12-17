package entities

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

type ICommentsRepository interface {
	Create(comment *Comment) error
	GetAllFromAPost(postId string) (*[]CommentFromPersistence, error)
	UpdateContent(dto *Comment) error
	Delete(id string) error
}
