package likes

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ILikesRepository interface {
	LikePost(data *LikePostInPersistence) error
	DislikePost(data *LikePostInPersistence) error
}

type LikesRepository struct {
	DB *sqlx.DB
}

func (r *LikesRepository) LikePost(data *LikePostInPersistence) error {
	query := `INSERT INTO likes(id, user_id, post_id) VALUES(:id, :user_id, :post_id)`

	_, err := r.DB.NamedExec(query, &data)
	if err != nil {
		return err
	}

	return nil
}

func (r *LikesRepository) DislikePost(data *LikePostInPersistence) error {
	var exists bool

	existenceQuery := `SELECT EXISTS(
		SELECT id FROM likes 
		WHERE user_id=$1 
		AND post_id=$2
		) as like_exists;
	`

	err := r.DB.Get(&exists, existenceQuery, data.UserID, data.PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		unknownError := fmt.Errorf("unknown error while querying likes: %v", err.Error())

		return unknownError
	}

	query := `DELETE FROM likes 
	WHERE WHERE user_id=$1 
	AND post_id=$2`

	_, err = r.DB.Exec(query, data.UserID, data.PostID)
	if err != nil {
		unknownError := fmt.Errorf("unknown error while querying likes: %v", err.Error())

		return unknownError
	}

	return nil
}
