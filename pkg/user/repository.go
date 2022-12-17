package user

import (
	"database/sql"
	"log"

	"github.com/iugstav/colatech-api/internal/entities"
	"github.com/jmoiron/sqlx"
)

type UsersRepository struct {
	DB *sqlx.DB
}

func GenerateNewUsersRepository(db *sqlx.DB) entities.IUsersRepository {
	return &UsersRepository{DB: db}
}

func (r *UsersRepository) Create(user *entities.User) error {
	query := `INSERT INTO readers(id, username, first_name, last_name, email, password, image_url, created_at)
			  VALUES(:id, :username, :first_name, :last_name, :email, :password, '', :created_at)`

	_, err := r.DB.NamedExec(query, user)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) GetById(id string) (*entities.User, error) {
	var user entities.User

	err := r.DB.Get(&user, "SELECT * FROM readers WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) GetByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := r.DB.Get(&user, "SELECT * FROM readers WHERE email=$1", email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UsersRepository) UploadImage(dto *entities.UploadProfileImageToPersistence) error {
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
