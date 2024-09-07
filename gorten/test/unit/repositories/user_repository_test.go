package repositories_test

import (
	"context"
	"testing"

	"gorten/internal/gorten/models"
	"gorten/test/factories"
	"gorten/test/integration/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	userJohn := factories.UserFactory()
	userJane := factories.UserFactory(func(u *models.User) {
		u.Name = "Jane Doe"
	})

	expectedUsers := []models.User{*userJohn, *userJane}
	mockRepo := new(mocks.MockUserRepository)
	mockRepo.On("GetAll", ctx, 0, 10).Return(expectedUsers, nil)
	users, err := mockRepo.GetAll(ctx, 0, 10)

	require.NoError(t, err)
	assert.Len(t, users, 2)
	assert.Equal(t, userJohn.Name, users[0].Name)
	assert.Equal(t, userJane.Name, users[1].Name)
	mockRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	ctx := context.Background()
	newUser := factories.UserFactory()
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("GetByID", ctx, newUser.UserID).Return(newUser, nil)
	user, err := mockRepo.GetByID(ctx, newUser.UserID)

	require.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, newUser.Name, user.Name)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	ctx := context.Background()
	newUser := factories.UserFactory()
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("Create", ctx, newUser).Return(nil)
	err := mockRepo.Create(ctx, newUser)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	user := factories.UserFactory()
	mockRepo := new(mocks.MockUserRepository)

	mockRepo.On("Update", ctx, user).Return(nil)
	err := mockRepo.Update(ctx, user)

	require.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
