package services

import (
	"fmt"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func UserServiceInit(repo repositories.UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

func (s *UserService) List() ([]models.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users: %w", err)
	}
	return users, nil
}

func (s *UserService) GetByID(id string) (*models.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %w", err)
	}
	return user, nil
}

func (s *UserService) Create(user *models.User) error {
	err := s.userRepo.Create(user)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (s *UserService) UpdateByID(id string, user *models.User) error {
	existingUser, err := s.userRepo.GetByID(id)
	if err != nil {
		return fmt.Errorf("failed to retrieve user with ID %s: %w", id, err)
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := s.userRepo.Update(existingUser); err != nil {
		return fmt.Errorf("failed to update user with ID %s: %w", id, err)
	}

	return nil
}
