package mocks

import (
	"bytes"

	"github.com/sirupsen/logrus"
)

type MockLogger struct {
	*logrus.Logger
	Buffer *bytes.Buffer
}

func NewMockLogger() *MockLogger {
	buffer := new(bytes.Buffer)
	logger := logrus.New()
	logger.SetOutput(buffer)
	return &MockLogger{
		Logger: logger,
		Buffer: buffer,
	}
}

func (m *MockLogger) Log(level logrus.Level, args ...interface{}) {
	m.Logger.Log(level, args...)
}

func (m *MockLogger) Logf(level logrus.Level, format string, args ...interface{}) {
	m.Logger.Logf(level, format, args...)
}

func (m *MockLogger) Logln(level logrus.Level, args ...interface{}) {
	m.Logger.Logln(level, args...)
}
