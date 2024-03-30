package service

import (
	"context"
	"go-rest-api/internal/app/model"
	"go-rest-api/internal/app/repository"
)

// PostService defines the interface for post service operations
type PostService interface {
	CreatePost(ctx context.Context, post *model.Post) error
	GetPostByID(ctx context.Context, id int64) (*model.Post, error)
}

// postServiceImpl implements PostService with a repository layer.
type postServiceImpl struct {
	postRepo repository.PostRepository
}

// NewPostService creates a new instance of PostService
func NewPostService(postRepo repository.PostRepository) PostService {
	return &postServiceImpl{postRepo: postRepo}
}

// CreatePost handles the creation of a new post
func (s *postServiceImpl) CreatePost(ctx context.Context, post *model.Post) error {
	// Here, we can add some business logic, like validation or preprocessing, before creating the post
	return s.postRepo.Create(ctx, post)
}

// GetPostByID handles fetching a post by its ID
func (s *postServiceImpl) GetPostByID(ctx context.Context, id int64) (*model.Post, error) {
	// Here, we can add some business logic, like checking if the post exists, before fetching the post
	return s.postRepo.FindByID(ctx, id)
}
