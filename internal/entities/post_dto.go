package entities

import (
	"os"
	"time"
)

type ResumedPost struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	CoverImageURL string    `json:"cover_image_url"`
	Intro         string    `json:"intro"`
	CategoryID    string    `json:"category_id"`
	PublishedAt   time.Time `json:"published_at"`
}

type GetPostByIdServiceResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Slug          string    `json:"slug"`
	CoverImageURL string    `json:"cover_image_url"`
	Content       string    `json:"content"`
	Category      Category  `json:"category"`
	PublishedAt   time.Time `json:"published_at"`

	Comments []CommentFromPersistence `json:"comments"`
}

type GetPostByIdFromRepository struct {
	ID            string    `json:"id" db:"id"`
	Title         string    `json:"title" db:"title"`
	Slug          string    `json:"slug" db:"slug"`
	CoverImageURL string    `json:"cover_image_url" db:"cover_image_url"`
	Content       string    `json:"content" db:"content"`
	CategoryID    string    `json:"category_id" db:"category_id"`
	CategoryName  string    `json:"category_name" db:"category_name"`
	PublishedAt   time.Time `json:"published_at" db:"published_at"`
}

type CreatePostFromEndpoint struct {
	Title       string `form:"title" binding:"required,min=8,max=255"`
	Slug        string `form:"slug" binding:"required,contains=-"`
	Intro       string `form:"intro" binding:"required"`
	Content     string `form:"content" binding:"required"`
	CategoryID  string `form:"category_id" binding:"required"`
	PublishedAt string `form:"published_at" binding:"required"`
}

type CreatePostServiceRequest struct {
	// request data from endpoint
	ID          string
	Title       string
	Slug        string
	Intro       string
	Content     string
	CategoryID  string
	PublishedAt string
}

type UploadPostCoverImageFromEndpoint struct {
	PostID string `json:"post_id" binding:"required"`
}

type UploadPostCoverImageRequest struct {
	File                 *os.File
	ID                   string
	NameToUpload         string
	CoverImageStaticPath string
}

type UploadPostCoverImageInPersistence struct {
	ID            string
	CoverImageURL string
}

type UpdatePostContentDTO struct {
	ID         string `json:"id" binding:"required,len=36"`
	NewContent string `json:"new_content" binding:"required"`
}

type LikePostDTO struct {
	UserID string `json:"user_id" binding:"required"`
	PostID string `json:"post_id" binding:"required"`
}
