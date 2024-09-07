package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetAll(ctx context.Context, skip, limit int) ([]models.User, error) {
	args := m.Called(ctx, skip, limit)
	users, _ := args.Get(0).([]models.User)
	err, _ := args.Get(1).(error)
	return users, err
}

func (m *MockUserRepository) GetByID(ctx context.Context, userID string) (*models.User, error) {
	args := m.Called(ctx, userID)
	user, _ := args.Get(0).(*models.User)
	err, _ := args.Get(1).(error)
	return user, err
}

func (m *MockUserRepository) Create(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockUserRepository) Update(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockUserRepository.Update: %w", err)
	}
	return nil
}
