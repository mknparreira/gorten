package services

import (
	"context"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/repositories"
	pkgerr "gorten/pkg/errors"
	"gorten/pkg/logs"
	"gorten/pkg/utils"
)

type UserServiceImpl interface {
	List(ctx context.Context, skip, limit int, sort string) ([]models.User, error)
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

func (s *UserService) List(ctx context.Context, skip, limit int, sort string) ([]models.User, error) {
	users, err := s.userRepo.GetAll(ctx, skip, limit, sort)
	if err != nil {
		logs.Logger.Printf("Error on UserService::List. Reason: %v", err)
		return nil, pkgerr.ErrInternalServerError
	}
	return users, nil
}

func (s *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on UserService::GetByID. Reason: %v", err)
		return nil, pkgerr.ErrUserNotFound
	}
	return user, nil
}

func (s *UserService) Create(ctx context.Context, user *models.User) error {
	userID, _ := utils.GenerateUUID()
	user.UserID = userID

	err := s.userRepo.Create(ctx, user)
	if err != nil {
		logs.Logger.Printf("Error on UserService::Create. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}

func (s *UserService) UpdateByID(ctx context.Context, id string, user *models.User) error {
	existingUser, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		logs.Logger.Printf("Error on retrieve GetByID into UserService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrUserNotFound
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password

	if err := s.userRepo.Update(ctx, existingUser); err != nil {
		logs.Logger.Printf("Error on UserService::UpdateByID. Reason: %v", err)
		return pkgerr.ErrInternalServerError
	}
	return nil
}
