package mocks

import (
	"context"
	"fmt"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoClientMock struct {
	mock.Mock
}

func (m *MongoClientMock) Ping(ctx context.Context, rp *readpref.ReadPref) error {
	args := m.Called(ctx, rp)
	if err := args.Error(0); err != nil {
		return fmt.Errorf("MongoClientMock.Ping: %w", err)
	}
	return nil
}

func (m *MongoClientMock) Disconnect(ctx context.Context) error {
	args := m.Called(ctx)
	if err := args.Error(0); err != nil {
		return fmt.Errorf("MongoClientMock.Disconnect: %w", err)
	}
	return nil
}

func (m *MongoClientMock) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	args := m.Called(name, opts)
	return args.Get(0).(*mongo.Database)
}
