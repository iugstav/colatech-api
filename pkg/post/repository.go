package post

import (
	"github.com/iugstav/colatech-api/pkg/comment"
	"github.com/jmoiron/sqlx"
)

type IPostsRepository interface {
	Create(post *Post) error
	GetAll() ([]*Post, error)
	GetAllMinified() ([]*ResumedPost, error)
	GetById(id string) (*PostFromPersistence, error)
	UpdateContent(dto *UpdatePostContentDTO) error
	UploadImage(dto *UploadPostCoverImageInPersistence) error
	Delete(id string) error
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

func (r *PostsRepository) GetById(id string) (*PostFromPersistence, error) {
	var post PostFromPersistence
	var postComments []comment.Comment

	// TODO: analisar melhor query a ser utilizada pro fetch de posts. acho que a de comentários tá boa

	postsQuery := `SELECT p.id, p.title, p.slug, p.content, p.category_id, p.published_at, p.cover_image_url, c.name AS category_name 
			FROM posts p 
			LEFT JOIN categories c ON p.category_id=c.id 
			WHERE p.id=$1`

	commentsQuery := `SELECT p.id, p.reader_id, p.post_id, 
				p.parent_id, p.content, p.published_at, 
				r.first_name, r.last_name, r.image_url FROM post_comments p 
				LEFT JOIN readers r ON p.reader_id=r.id
				WHERE p.post_id=$1`

	err := r.DB.Get(&post, postsQuery, id)
	if err != nil {
		return nil, err
	}

	err = r.DB.Select(&postComments, commentsQuery, id)
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

func (r *PostsRepository) Delete(id string) error {
	_, err := r.DB.Exec(`DELETE FROM posts WHERE id=$1`, id)
	if err != nil {
		return err
	}

	return nil
}
