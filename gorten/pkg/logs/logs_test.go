package logs

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestInitLogger(t *testing.T) {
	logger := InitLogger()

	assert.NotNil(t, logger, "Logger should not be nil")
	assert.IsType(t, &logrus.Logger{}, logger, "Logger should be of type *logrus.Logger")
	assert.Equal(t, logrus.DebugLevel, logger.GetLevel(), "Logger level should be DebugLevel")

	formatter, ok := logger.Formatter.(*logrus.JSONFormatter)
	assert.True(t, ok, "Logger formatter should be of type *logrus.JSONFormatter")
	assert.True(t, formatter.PrettyPrint, "Formatter PrettyPrint should be true")
	assert.True(t, formatter.DisableHTMLEscape, "Formatter DisableHTMLEscape should be true")
	assert.Equal(t, "@timestamp", formatter.FieldMap[logrus.FieldKeyTime], "FieldMap time key should be '@timestamp'")
	assert.Equal(t, "severity", formatter.FieldMap[logrus.FieldKeyLevel], "FieldMap level key should be 'severity'")
	assert.Equal(t, "message", formatter.FieldMap[logrus.FieldKeyMsg], "FieldMap msg key should be 'message'")
}
