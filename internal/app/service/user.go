package service

import (
	"context"
	"go-rest-api/internal/app/model"
	"go-rest-api/internal/app/repository"
)

// UserService defines the interface for user service operations
type UserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUserByID(ctx context.Context, id int64) (*model.User, error)

	// Add more methods here
}

// userServiceImpl implements UserService with a repository layer.
type userServiceImpl struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{userRepo: userRepo}
}

// CreateUser handles the creation of a new user
func (s *userServiceImpl) CreateUser(ctx context.Context, user *model.User) error {
	// Here, we can add some business logic, like validation or preprocessing, before creating the user
	return s.userRepo.Create(ctx, user)
}

// GetUserByID handles fetching a user by their ID
func (s *userServiceImpl) GetUserByID(ctx context.Context, id int64) (*model.User, error) {
	// Here, we can add some business logic, like checking if the user exists, before fetching the user
	return s.userRepo.FindByID(ctx, id)
}
