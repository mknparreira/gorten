package services

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/repositories"
)

type UserServiceImpl interface {
	List(ctx context.Context) ([]models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	Create(ctx context.Context, user *models.User) error
	UpdateByID(ctx context.Context, id string, user *models.User) error
}

type UserService struct {
	userRepo repositories.UserRepositoryImpl
}

func UserServiceInit(repo repositories.UserRepositoryImpl) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) List(ctx context.Context) ([]models.User, error) {
	users, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	err := s.userRepo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *UserService) UpdateByID(ctx context.Context, id string, user *models.User) error {
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to retrieve user with ID %s: %w", id, err)
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		return fmt.Errorf("failed to update user with ID %s: %w", id, err)
	}

	return nil
}
