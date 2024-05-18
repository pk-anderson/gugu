package post

import (
	"database/sql"
	"errors"
	"fmt"
	post "gugu/interfaces/post"
	"gugu/repositories/postRepository"
	"gugu/utils"
)

type PostService interface {
	CreatePost(userId, content string) (string, error)
}

type postService struct {
	DB *sql.DB
}

func (s *postService) CreatePost(userId, content string) (string, error) {
	if len(content) > 4 {
		return "", errors.New("content must have a maximum of 4 characters")
	}
	fmt.Printf("1")
	rep := postRepository.NewRepository(s.DB)
	postId := utils.GenerateUUID()
	post := &post.Post{
		PostID:  postId,
		UserID:  userId,
		Content: content,
	}
	fmt.Printf("2")
	err := rep.InsertPost(post)
	if err != nil {
		return "", err
	}
	fmt.Printf("3")
	return postId, nil
}

func NewService(DB *sql.DB) PostService {
	return &postService{
		DB: DB,
	}
}
