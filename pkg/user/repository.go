package user

import (
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

type IUsersRepository interface {
	Create(user *User) error
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	UploadImage(dto *UploadProfileImageToPersistence) error
	Exists(id string) bool
}

type UsersRepository struct {
	DB *sqlx.DB
}

func GenerateNewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

func (r *UsersRepository) Create(user *User) error {
	query := `INSERT INTO readers(id, username, first_name, last_name, email, password, image_url, created_at)
			  VALUES(:id, :username, :first_name, :last_name, :email, :password, '', :created_at)`

	_, err := r.DB.NamedExec(query, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) GetById(id string) (*User, error) {
	var user User

	err := r.DB.Get(&user, "SELECT * FROM readers WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) GetByEmail(email string) (*User, error) {
	var user User

	err := r.DB.Get(&user, "SELECT * FROM readers WHERE email=$1", email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) UploadImage(dto *UploadProfileImageToPersistence) error {
	_, err := r.DB.Exec(`UPDATE readers SET image_url=$1 WHERE id=$2`, dto.ProfileImageURL, dto.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) Exists(id string) bool {
	var exists bool

	query := `SELECT EXISTS (SELECT id FROM readers where id=$1)`

	err := r.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			exists = false
		} else {
			log.Fatalf("Error while checking existence of user: %v", err.Error())
		}
	} else {
		exists = true
	}

	return exists
}
