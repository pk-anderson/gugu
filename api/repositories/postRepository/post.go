package postRepository

import (
	"database/sql"
	post "gugu/interfaces/post"
)

type PostRepository interface {
	InsertPost(post *post.Post) error
}

type postRepository struct {
	DB *sql.DB
}

func (r *postRepository) InsertPost(post *post.Post) error {
	query := `INSERT INTO tb_posts (post_id, user_id, content) 
	VALUES ($1, $2, $3)`

	_, err := r.DB.Exec(
		query,
		post.PostID,
		post.UserID,
		post.Content)
	return err
}

func NewRepository(DB *sql.DB) PostRepository {
	return &postRepository{
		DB: DB,
	}
}
