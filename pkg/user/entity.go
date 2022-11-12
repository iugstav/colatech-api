package user

import (
	"os"
	"time"
)

type User struct {
	ID        string    `json:"id" db:"id"`
	UserName  string    `json:"username" db:"username"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	ImageURL  string    `json:"image_url" db:"image_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type CreateUserServiceRequest struct {
	ID              string
	UserName        string `json:"username" binding:"required,min=4,max=24"`
	FirstName       string `json:"first_name" binding:"required,min=1,max=24"`
	LastName        string `json:"last_name" binding:"required,min=1,max=24"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6,max=32"`
	ImageStaticPath string
	CreatedAt       string `json:"created_at"`
}

type CreateUserFromEndpoint struct {
	UserName  string `form:"username" binding:"required,min=4,max=24"`
	FirstName string `form:"first_name" binding:"required,min=1,max=24"`
	LastName  string `form:"last_name" binding:"required,min=1,max=24"`
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required,min=6,max=32"`
	CreatedAt string `form:"created_at"`
}

type AuthenticateUserServiceRequest struct {
	Email    string `json:"email" db:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=32"`
}

type AuthenticateUserResponse struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

type GetUserByIdResponse struct {
	UserName  string    `json:"username" db:"user_name"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UploadProfileImageFromEndpoint struct {
	UserID string `json:"user_id" binding:"required"`
}

type UploadProfileImageRequest struct {
	File                   *os.File
	ID                     string
	NameToUpload           string
	ProfileImageStaticPath string
}

type UploadProfileImageToPersistence struct {
	ID              string
	ProfileImageURL string
}

type CreateFirebaseUserData struct {
	ID            string
	Email         string
	ImageURL      string
	BrutePassword string
}
