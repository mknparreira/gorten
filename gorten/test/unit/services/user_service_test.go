package services_test

import (
	"context"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/services"
	"gorten/pkg/errors"
	"gorten/test/integration/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserService_List(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	expectedUsers := []models.User{{UserID: "1", Name: "John Doe", Email: "john.doe@gmail.com"}}
	mockRepo.On("GetAll", ctx).Return(expectedUsers, nil)

	users, err := service.List(ctx)

	require.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetByID(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	expectedUser := &models.User{UserID: "1", Name: "John Doe"}
	mockRepo.On("GetByID", ctx, "1").Return(expectedUser, nil)
	user, err := service.GetByID(ctx, "1")

	require.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Create(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	newUser := &models.User{UserID: "1", Name: "John Doe"}
	mockRepo.On("Create", ctx, newUser).Return(nil)
	err := service.Create(ctx, newUser)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateByID(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	existingUser := &models.User{UserID: "1", Name: "John Doe"}
	updatedUser := &models.User{Name: "Jane Doe"}

	mockRepo.On("GetByID", ctx, "1").Return(existingUser, nil)
	mockRepo.On("Update", ctx, existingUser).Return(nil)

	err := service.UpdateByID(ctx, "1", updatedUser)

	require.NoError(t, err)
	assert.Equal(t, updatedUser.Name, existingUser.Name)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateByID_UserNotFound(t *testing.T) {
	ctx := context.TODO()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	mockRepo.On("GetByID", ctx, "1").Return(nil, errors.ErrUserNotFound)
	err := service.UpdateByID(ctx, "1", &models.User{Name: "Jane Doe"})

	require.Error(t, err)
	mockRepo.AssertExpectations(t)
}
