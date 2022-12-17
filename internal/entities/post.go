package entities

import (
	"time"
)

type Post struct {
	ID            string    `json:"id" db:"id"`
	Title         string    `json:"title" db:"title"`
	Slug          string    `json:"slug" db:"slug"`
	CoverImageURL string    `json:"cover_image_url" db:"cover_image_url"`
	Intro         string    `json:"intro" db:"intro"`
	Content       string    `json:"content" db:"content"`
	CategoryID    string    `json:"category_id" db:"category_id"`
	PublishedAt   time.Time `json:"published_at" db:"published_at"`
}

type IPostsRepository interface {
	Create(post *Post) error
	GetAll() ([]*Post, error)
	GetAllMinified() ([]*ResumedPost, error)
	GetById(id string) (*GetPostByIdFromRepository, error)
	UpdateContent(dto *UpdatePostContentDTO) error
	UploadImage(dto *UploadPostCoverImageInPersistence) error
	LikePost(data *LikePostInPersistence) error
	Delete(id string) error
	Exists(id string) bool
	BothUserAndPostExists(userId string, postId string) (bool, bool)
}
