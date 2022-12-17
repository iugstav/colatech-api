package entities

import (
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

type IUsersRepository interface {
	Create(user *User) error
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	UploadImage(dto *UploadProfileImageToPersistence) error
	Exists(id string) bool
}
