package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go-rest-api/internal/app/model"
)

// UserRepository defines the interface for user data operations
type UserRepository interface {
	FindByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, u *model.User) error

	// Add more methods as needed
}

// userRepositoryImpl implements the UserRepository with a Postgres database
type userRepositoryImpl struct {
	db *sql.DB
}

// NewUserRepository creates a new instance of UserRepository
// Constructor function
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

// FindByID retrieves a user by ID
func (repo *userRepositoryImpl) FindByID(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	query := `SELECT id, name, email FROM users WHERE id = $1`
	if err := repo.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, fmt.Errorf("USER NOT FOUND: %w", err)
	}

	return &user, nil
}

// Create inserts a new user
func (repo *userRepositoryImpl) Create(ctx context.Context, u *model.User) error {
	query := `INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`
	if err := repo.db.QueryRowContext(ctx, query, u.Name, u.Email).Scan(&u.ID); err != nil {
		return fmt.Errorf("FAILED TO CREATE USER: %w", err)
	}

	return nil
}
