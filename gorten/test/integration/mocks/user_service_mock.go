package mocks

import (
	"context"
	"fmt"
	"gorten/internal/gorten/models"

	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) List(ctx context.Context, skip, limit int) ([]models.User, error) {
	args := m.Called(ctx, skip, limit)

	users, _ := args.Get(0).([]models.User)
	err, _ := args.Get(1).(error)
	return users, err
}

func (m *MockUserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)
	user, _ := args.Get(0).(*models.User)
	err, _ := args.Get(1).(error)
	return user, err
}

func (m *MockUserService) Create(ctx context.Context, user *models.User) error {
	args := m.Called(ctx, user)
	if err, ok := args.Get(0).(error); ok {
		return err
	}
	return nil
}

func (m *MockUserService) UpdateByID(ctx context.Context, id string, user *models.User) error {
	args := m.Called(ctx, id, user)

	if err := args.Error(0); err != nil {
		return fmt.Errorf("MockUserService.Update: %w", err)
	}
	return nil
}
