package comment

import "github.com/jmoiron/sqlx"

type ICommentsRepository interface {
	Create(comment *Comment) error
	GetAllFromAPost(postId string) (*[]CommentFromPersistence, error)
	UpdateContent(dto *UpdateCommentContentDTO) error
	Delete(id string) error
}

type CommentsRepository struct {
	DB *sqlx.DB
}

func GenerateNewCommentsRepository(db *sqlx.DB) *CommentsRepository {
	return &CommentsRepository{DB: db}
}

func (r *CommentsRepository) Create(comment *Comment) error {
	query := `INSERT INTO post_comments(id, reader_id, post_id, parent_id, content, published_at)
			  VALUES(:id, :reader_id, :post_id, :parent_id, :content, :published_at)`

	_, err := r.DB.NamedExec(query, comment)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentsRepository) GetAllFromAPost(postId string) (*[]CommentFromPersistence, error) {
	var comments *[]CommentFromPersistence

	err := r.DB.Select(
		&comments,
		`SELECT c.id, c.reader_id, c.post_id, 
		 c.parent_id, c.content, c.published_at, 
		 r.first_name AS reader_first_name, r.last_name AS reader_last_name, r.image_url FROM post_comments c 
		 LEFT JOIN readers r ON p.reader_id=r.id
		 WHERE p.post_id=$1`,
		postId)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (r *CommentsRepository) UpdateContent(dto *UpdateCommentContentDTO) error {
	_, err := r.DB.Exec("UPDATE post_comments SET content=$1 WHERE id=$2", dto.ID, dto.NewContent)
	if err != nil {
		return err
	}

	return nil
}

func (r *CommentsRepository) Delete(id string) error {
	_, err := r.DB.Exec("DELETE FROM post_comments WHERE id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
