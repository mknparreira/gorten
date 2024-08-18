package db_test

import (
	"context"
	"gorten/internal/gorten/db"
	"gorten/pkg/errors"
	"gorten/test/integration/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	m.Run()
}

func setupConnection(mockClient *mocks.MongoClientMock) (context.Context, context.CancelFunc) {
	//Replace the real MongoClient for mock interface
	db.MongoClient = mockClient
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	return ctx, cancel
}

func TestConnect(t *testing.T) {
	mockClient := new(mocks.MongoClientMock)

	//Setting up mock for the Ping and Disconnect methods
	mockClient.On("Ping", mock.Anything, mock.Anything).Return(nil)
	mockClient.On("Disconnect", mock.Anything).Return(nil).Once()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db.MongoClient = mockClient

	err := db.MongoClient.Ping(ctx, nil)

	require.NoError(t, err, "Failed to ping MongoDB")
	assert.NotNil(t, db.MongoClient, "MongoClient should not be nil")

	db.Disconnect(ctx, cancel)

	mockClient.AssertExpectations(t)
}

func TestDisconnect(t *testing.T) {
	mockClient := new(mocks.MongoClientMock)
	ctx, cancel := setupConnection(mockClient)
	defer db.Disconnect(ctx, cancel)

	mockClient.On("Disconnect", mock.Anything).Return(nil)

	db.Disconnect(ctx, cancel)

	mockClient.On("Ping", mock.Anything, mock.Anything).Return(errors.ErrClientDisconnected)

	err := db.MongoClient.Ping(ctx, nil)

	require.Error(t, err, "Expected error after disconnecting")
	assert.Contains(t, err.Error(), "client is disconnected", "Expected client disconnected error, but got a different error")
	mockClient.AssertExpectations(t)
}

func TestConnectTimeout(t *testing.T) {
	mockClient := new(mocks.MongoClientMock)
	//Initializing the context with a very short timeout to simulate timeout
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	//Setting up mock to simulate a timeout error when trying to ping
	mockClient.On("Ping", mock.Anything, mock.Anything).Return(errors.ErrContextDeadlineExceeded)
	db.MongoClient = mockClient
	err := db.MongoClient.Ping(ctx, nil)

	require.Error(t, err, "Expected error due to timeout")
	assert.Contains(t, err.Error(), "context deadline exceeded", "Expected context deadline exceeded error, but got a different error")
	mockClient.AssertExpectations(t)
}
