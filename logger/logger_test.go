package logger

import (
	"context"
	"os"
	"testing"
	"user-search-service/constants/api_constants"

	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// setupTestLogger sets up a test logger without config dependencies
func setupTestLogger() {
	logrusVar := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: joonix.NewFormatter(),
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}
	Log = &Logger{logrusVar}
}

func TestSetupLogger(t *testing.T) {
	// Test that we can initialize the global Log variable
	setupTestLogger()

	assert.NotNil(t, Log)
	assert.NotNil(t, Log.Logger)
	assert.Equal(t, logrus.InfoLevel, Log.Level)
}

func TestBuildContext(t *testing.T) {
	contextStr := "test_context"
	fields := BuildContext(contextStr)

	assert.Equal(t, contextStr, fields["context"])
	assert.NotEmpty(t, fields["request_id"])
	assert.IsType(t, "", fields["request_id"])
}

func TestGetStringFromContext(t *testing.T) {
	ctx := context.Background()
	value := "test_value"

	// Test with string value
	ctx = context.WithValue(ctx, api_constants.RequestId, value)
	result := GetStringFromContext(ctx, api_constants.RequestId)
	assert.Equal(t, value, result)

	// Test with non-string value
	ctx = context.WithValue(ctx, api_constants.UserId, 123)
	result = GetStringFromContext(ctx, api_constants.UserId)
	assert.Equal(t, "", result)

	// Test with non-existent key (using a different constant)
	result = GetStringFromContext(ctx, api_constants.DeviceId)
	assert.Equal(t, "", result)
}

func TestGetLoggerFromContext(t *testing.T) {
	// Setup logger first
	setupTestLogger()

	// Test with nil context
	logger := GetLoggerFromContext(context.TODO())
	assert.NotNil(t, logger)
	assert.IsType(t, &logrus.Entry{}, logger)

	// Test with context containing user data
	ctx := context.Background()
	ctx = context.WithValue(ctx, api_constants.RequestId, "req-123")
	ctx = context.WithValue(ctx, api_constants.UserId, "user-456")
	ctx = context.WithValue(ctx, api_constants.UserSessionId, "session-789")
	ctx = context.WithValue(ctx, api_constants.DeviceId, "device-abc")

	logger = GetLoggerFromContext(ctx)
	assert.NotNil(t, logger)

	// Check that the fields are properly set
	assert.Equal(t, "req-123", logger.Data[api_constants.RequestId])
	assert.Equal(t, "user-456", logger.Data[api_constants.UserId])
	assert.Equal(t, "session-789", logger.Data[api_constants.UserSessionId])
	assert.Equal(t, "device-abc", logger.Data[api_constants.DeviceId])
}

func TestLoggingFunctions(t *testing.T) {
	// Setup logger first
	setupTestLogger()

	ctx := context.Background()
	ctx = context.WithValue(ctx, api_constants.RequestId, "test-req")

	// These tests mainly ensure the functions don't panic
	// In a real scenario, you might want to capture the output and verify it

	t.Run("Debug", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Debug(ctx, "test debug message")
			Debugf(ctx, "test debug message with format: %s", "param")
		})
	})

	t.Run("Info", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Info(ctx, "test info message")
			Infof(ctx, "test info message with format: %s", "param")
		})
	})

	t.Run("Warn", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Warn(ctx, "test warn message")
			Warnf(ctx, "test warn message with format: %s", "param")
		})
	})

	t.Run("Error", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Error(ctx, "test error message")
			Errorf(ctx, "test error message with format: %s", "param")
		})
	})
}

func TestLoggerWithNilContext(t *testing.T) {
	// Setup logger first
	setupTestLogger()

	// Test that logging functions handle nil context gracefully
	assert.NotPanics(t, func() {
		Debug(context.TODO(), "test with nil context")
		Info(context.TODO(), "test with nil context")
		Warn(context.TODO(), "test with nil context")
		Error(context.TODO(), "test with nil context")
	})
}

func TestLoggerError(t *testing.T) {
	err := &LoggerError{Error: assert.AnError}
	assert.Equal(t, assert.AnError, err.Error)
}

func TestPanicIfError(t *testing.T) {
	// Test with nil error (should not panic)
	assert.NotPanics(t, func() {
		panicIfError(nil)
	})

	// Test with actual error (should panic)
	assert.Panics(t, func() {
		panicIfError(assert.AnError)
	})
}
