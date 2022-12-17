package entities

type LikePostInPersistence struct {
	ID     string `db:"id"`
	UserID string `db:"user_id"`
	PostID string `db:"post_id"`
}

type ILikesRepository interface {
	LikePost(data *LikePostInPersistence) error
	DislikePost(data *LikePostInPersistence) error
}
