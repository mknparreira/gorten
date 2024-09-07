package services_test

import (
	"context"
	"gorten/internal/gorten/models"
	"gorten/internal/gorten/services"
	"gorten/pkg/errors"
	"gorten/test/factories"
	"gorten/test/integration/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const janeDoeName = "Jane Doe"

func TestUserService_List(t *testing.T) {
	ctx := context.Background()
	user := factories.UserFactory()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	expectedUsers := []models.User{*user}
	mockRepo.On("GetAll", ctx, 0, 10).Return(expectedUsers, nil)

	users, err := service.List(ctx, 0, 10)

	require.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetByID(t *testing.T) {
	ctx := context.Background()
	expectedUser := factories.UserFactory()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	mockRepo.On("GetByID", ctx, expectedUser.UserID).Return(expectedUser, nil)
	user, err := service.GetByID(ctx, expectedUser.UserID)

	require.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserService_Create(t *testing.T) {
	ctx := context.Background()
	newUser := factories.UserFactory()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	mockRepo.On("Create", ctx, newUser).Return(nil)
	err := service.Create(ctx, newUser)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateByID(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)

	existingUser := factories.UserFactory()
	updatedUser := factories.UserFactory(func(u *models.User) {
		u.Name = janeDoeName
	})

	mockRepo.On("GetByID", ctx, existingUser.UserID).Return(existingUser, nil)
	mockRepo.On("Update", ctx, existingUser).Return(nil)

	err := service.UpdateByID(ctx, existingUser.UserID, updatedUser)

	require.NoError(t, err)
	assert.Equal(t, updatedUser.Name, existingUser.Name)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateByID_UserNotFound(t *testing.T) {
	ctx := context.Background()
	mockRepo := new(mocks.MockUserRepository)
	service := services.UserServiceInit(mockRepo)
	user := factories.UserFactory()
	updatedUser := factories.UserFactory(func(u *models.User) {
		u.Name = janeDoeName
	})

	mockRepo.On("GetByID", ctx, user.UserID).Return(nil, errors.ErrUserNotFound)
	err := service.UpdateByID(ctx, user.UserID, updatedUser)

	require.Error(t, err)
	mockRepo.AssertExpectations(t)
}
