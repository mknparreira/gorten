package services_test

import (
	"gorten/pkg/logs"
	"gorten/test/integration/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	mockLogger := mocks.NewMockLogger()
	logs.Logger = mockLogger.Logger
	m.Run()
}
