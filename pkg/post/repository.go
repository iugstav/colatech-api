package post

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/iugstav/colatech-api/pkg/likes"
	"github.com/jmoiron/sqlx"
)

type IPostsRepository interface {
	Create(post *Post) error
	GetAll() ([]*Post, error)
	GetAllMinified() ([]*ResumedPost, error)
	GetById(id string) (*GetPostByIdFromRepository, error)
	UpdateContent(dto *UpdatePostContentDTO) error
	UploadImage(dto *UploadPostCoverImageInPersistence) error
	LikePost(data *likes.LikePostInPersistence) error
	Delete(id string) error
	Exists(id string) bool
	BothUserAndPostExists(userId string, postId string) (bool, bool)
}

type PostsRepository struct {
	DB *sqlx.DB
}

func GenerateNewPostsRepository(db *sqlx.DB) *PostsRepository {
	return &PostsRepository{DB: db}
}

func (r *PostsRepository) Create(post *Post) error {
	query := `INSERT INTO posts(id, title, slug, intro, content, category_id, published_at, cover_image_url)
	VALUES(:id, :title, :slug, :intro, :content, :category_id, :published_at, :cover_image_url)`

	_, err := r.DB.NamedExec(query, &post)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostsRepository) UploadImage(dto *UploadPostCoverImageInPersistence) error {
	query := `UPDATE posts SET cover_image_url=$1 WHERE id=$2`

	if _, err := r.DB.Exec(query, dto.CoverImageURL, dto.ID); err != nil {
		return err
	}

	return nil
}

func (r *PostsRepository) GetAll() ([]*Post, error) {
	posts := []*Post{}
	query := `SELECT * FROM posts ORDER BY published_at ASC`

	err := r.DB.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostsRepository) GetAllMinified() ([]*ResumedPost, error) {
	posts := []*ResumedPost{}
	query := `SELECT id, title, slug, cover_image_url, intro, category_id, published_at
	FROM posts ORDER BY published_at ASC`

	err := r.DB.Select(&posts, query)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *PostsRepository) GetById(id string) (*GetPostByIdFromRepository, error) {
	var post GetPostByIdFromRepository

	if exists := r.Exists(id); !exists {
		errorMsg := fmt.Errorf("GetById: post with provided id %s does not exists", id)

		return nil, errorMsg
	}

	postsQuery := `SELECT p.id, p.title, p.slug, p.content, p.category_id, p.published_at, p.cover_image_url, c.name AS category_name 
			FROM posts p 
			LEFT JOIN categories c ON p.category_id=c.id 
			WHERE p.id=$1`

	err := r.DB.Get(&post, postsQuery, id)
	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *PostsRepository) UpdateContent(dto *UpdatePostContentDTO) error {
	_, err := r.DB.Exec(`UPDATE posts SET content=$1 WHERE id=$2`, dto.NewContent, dto.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostsRepository) LikePost(data *likes.LikePostInPersistence) error {
	query := `INSERT INTO likes(id, user_id, post_id) VALUES(:id, :user_id, :post_id)`

	_, err := r.DB.NamedExec(query, &data)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostsRepository) Delete(id string) error {
	_, err := r.DB.Exec(`DELETE FROM posts WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostsRepository) Exists(id string) bool {
	var exists bool

	query := `SELECT EXISTS (SELECT id FROM posts where id=$1)`

	err := r.DB.QueryRow(query, id).Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatalf("Error while checking existence of post: %v", err.Error())
		}
	}

	return true
}

func (r *PostsRepository) BothUserAndPostExists(userId string, postId string) (bool, bool) {
	var userExists bool
	var postExists bool

	userQuery := `SELECT EXISTS (SELECT id FROM readers where id=$1)`
	postQuery := `SELECT EXISTS (SELECT id FROM posts where id=$1)`

	err := r.DB.QueryRow(userQuery, userId).Scan(&userExists)
	if err != nil {
		if err == sql.ErrNoRows {
			userExists = false
		} else {
			log.Fatalf("Error while checking existence of user: %v", err.Error())
		}
	} else {
		userExists = true
	}

	err = r.DB.QueryRow(postQuery, postId).Scan(&postExists)
	if err != nil {
		if err == sql.ErrNoRows {
			postExists = false
		} else {
			log.Fatalf("Error while checking existence of post: %v", err.Error())
		}
	} else {
		postExists = true
	}

	return postExists, userExists
}
