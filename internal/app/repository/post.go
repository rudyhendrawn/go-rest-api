package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-rest-api/internal/app/model"
)

// PostRepository defines the interface for post data operations
type PostRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Post, error)
	Create(ctx context.Context, p *model.Post) error
	GetAllPosts(ctx context.Context) ([]*model.Post, error)
	// Add more methods as needed
}

// PostRepositoryImpl implements the PostRepository with a Postgres database
type postRepositoryImpl struct {
	db *sql.DB
}

// NewPostRepository creates a new instance of PostRepository
// Constructor function
func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepositoryImpl{db: db}
}

// NewPostRepository creates a new instance of PostRepository
func (repo *postRepositoryImpl) FindByID(ctx context.Context, id int64) (*model.Post, error) {
	var post model.Post
	query := `SELECT id, user_id, title, content FROM posts WHERE id = $1`
	if err := repo.db.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.UserID, &post.Title, &post.Content); err != nil {
		return nil, fmt.Errorf("POST NOT FOUND: %w", err)
	}

	return &post, nil
}

// Create inserts a new post
func (repo *postRepositoryImpl) Create(ctx context.Context, p *model.Post) error {
	query := `INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3) RETURNING id`
	if err := repo.db.QueryRowContext(ctx, query, p.UserID, p.Title, p.Content).Scan(&p.ID); err != nil {
		return fmt.Errorf("FAILED TO CREATE POST: %w", err)
	}

	return nil
}

// GetAllPosts retrieves all posts
func (repo *postRepositoryImpl) GetAllPosts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post
	query := `SELECT id, user_id, title, content FROM posts`
	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("FAILED TO GET ALL POSTS: %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post.ID, &post.UserID, &post.Title, &post.Content); err != nil {
			return nil, fmt.Errorf("FAILED TO SCAN POST: %w", err)
		}

		posts = append(posts, &post)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("FAILED TO GET ALL POSTS: %w", err)
	}

	return posts, nil
}
