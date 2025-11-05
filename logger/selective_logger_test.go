package logger

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestAddEnabledLoggingComponentsToContext(t *testing.T) {
	enabledComponents := []string{"search_flow", "ranking_service", "ads_integration"}

	t.Run("with valid context", func(t *testing.T) {
		ctx := context.Background()
		resultCtx := AddEnabledLoggingComponentsToContext(ctx, enabledComponents)

		assert.NotNil(t, resultCtx)

		// Verify the components are stored in context
		stored, ok := resultCtx.Value(LoggingEnabledComponentsKey).([]string)
		assert.True(t, ok)
		assert.Equal(t, enabledComponents, stored)
	})

	t.Run("with nil context", func(t *testing.T) {
		resultCtx := AddEnabledLoggingComponentsToContext(context.TODO(), enabledComponents)

		assert.NotNil(t, resultCtx)

		// Verify the components are stored in context
		stored, ok := resultCtx.Value(LoggingEnabledComponentsKey).([]string)
		assert.True(t, ok)
		assert.Equal(t, enabledComponents, stored)
	})

	t.Run("with empty components list", func(t *testing.T) {
		ctx := context.Background()
		resultCtx := AddEnabledLoggingComponentsToContext(ctx, []string{})

		assert.NotNil(t, resultCtx)

		stored, ok := resultCtx.Value(LoggingEnabledComponentsKey).([]string)
		assert.True(t, ok)
		assert.Empty(t, stored)
	})
}

func TestIsComponentEnabled(t *testing.T) {
	enabledComponents := []string{"search_flow", "ranking_service", "ads_integration"}

	t.Run("component is enabled", func(t *testing.T) {
		ctx := AddEnabledLoggingComponentsToContext(context.Background(), enabledComponents)

		assert.True(t, isComponentEnabled(ctx, "search_flow"))
		assert.True(t, isComponentEnabled(ctx, "ranking_service"))
		assert.True(t, isComponentEnabled(ctx, "ads_integration"))
	})

	t.Run("component is not enabled", func(t *testing.T) {
		ctx := AddEnabledLoggingComponentsToContext(context.Background(), enabledComponents)

		assert.False(t, isComponentEnabled(ctx, "disabled_component"))
		assert.False(t, isComponentEnabled(ctx, "another_component"))
	})

	t.Run("case insensitive matching", func(t *testing.T) {
		ctx := AddEnabledLoggingComponentsToContext(context.Background(), []string{"Search_Flow", "RANKING_SERVICE"})

		assert.True(t, isComponentEnabled(ctx, "search_flow"))
		assert.True(t, isComponentEnabled(ctx, "Search_Flow"))
		assert.True(t, isComponentEnabled(ctx, "ranking_service"))
		assert.True(t, isComponentEnabled(ctx, "RANKING_SERVICE"))
	})

	t.Run("exact matching without trimming", func(t *testing.T) {
		// Components with spaces should not match components without spaces
		ctx := AddEnabledLoggingComponentsToContext(context.Background(), []string{" search_flow ", "ranking_service"})

		assert.True(t, isComponentEnabled(ctx, " search_flow "))   // Exact match
		assert.False(t, isComponentEnabled(ctx, "search_flow"))    // No trimming
		assert.True(t, isComponentEnabled(ctx, "ranking_service")) // Exact match
	})

	t.Run("nil context", func(t *testing.T) {
		assert.False(t, isComponentEnabled(context.TODO(), "any_component"))
	})

	t.Run("context without enabled components", func(t *testing.T) {
		ctx := context.Background()
		assert.False(t, isComponentEnabled(ctx, "any_component"))
	})

	t.Run("context with wrong type", func(t *testing.T) {
		ctx := context.WithValue(context.Background(), LoggingEnabledComponentsKey, "not_a_slice")
		assert.False(t, isComponentEnabled(ctx, "any_component"))
	})

	t.Run("empty components list", func(t *testing.T) {
		ctx := AddEnabledLoggingComponentsToContext(context.Background(), []string{})
		assert.False(t, isComponentEnabled(ctx, "any_component"))
	})
}

func TestGetComponentLogger(t *testing.T) {
	// Setup main logger first
	setupTestLogger()

	ctx := context.Background()
	componentName := "test_component"

	logger := getComponentLogger(ctx, componentName)

	assert.NotNil(t, logger)
	assert.IsType(t, &logrus.Entry{}, logger)
	assert.Equal(t, componentName, logger.Data[Component])
}

func TestSelectiveLoggingFunctions(t *testing.T) {
	// Setup main logger first
	setupTestLogger()

	enabledComponents := []string{"enabled_component"}
	ctx := AddEnabledLoggingComponentsToContext(context.Background(), enabledComponents)

	t.Run("enabled component logging", func(t *testing.T) {
		// These should not panic and should execute the logging
		assert.NotPanics(t, func() {
			SelectiveDebug(ctx, "enabled_component", "debug message")
			SelectiveDebugf(ctx, "enabled_component", "debug message with param: %s", "value")
			SelectiveInfo(ctx, "enabled_component", "info message")
			SelectiveInfof(ctx, "enabled_component", "info message with param: %s", "value")
			SelectiveWarn(ctx, "enabled_component", "warn message")
			SelectiveWarnf(ctx, "enabled_component", "warn message with param: %s", "value")
			SelectiveError(ctx, "enabled_component", "error message")
			SelectiveErrorf(ctx, "enabled_component", "error message with param: %s", "value")
		})
	})

	t.Run("disabled component logging", func(t *testing.T) {
		// These should not panic but should not execute the logging
		assert.NotPanics(t, func() {
			SelectiveDebug(ctx, "disabled_component", "debug message")
			SelectiveDebugf(ctx, "disabled_component", "debug message with param: %s", "value")
			SelectiveInfo(ctx, "disabled_component", "info message")
			SelectiveInfof(ctx, "disabled_component", "info message with param: %s", "value")
			SelectiveWarn(ctx, "disabled_component", "warn message")
			SelectiveWarnf(ctx, "disabled_component", "warn message with param: %s", "value")
			SelectiveError(ctx, "disabled_component", "error message")
			SelectiveErrorf(ctx, "disabled_component", "error message with param: %s", "value")
		})
	})

	t.Run("nil context", func(t *testing.T) {
		// Should not panic with nil context
		assert.NotPanics(t, func() {
			SelectiveDebug(context.TODO(), "component", "message")
			SelectiveInfo(context.TODO(), "component", "message")
			SelectiveWarn(context.TODO(), "component", "message")
			SelectiveError(context.TODO(), "component", "message")
		})
	})
}

func TestSelectiveLoggingWithMultipleComponents(t *testing.T) {
	// Setup main logger first
	setupTestLogger()

	enabledComponents := []string{"search_flow", "ranking_service"}
	ctx := AddEnabledLoggingComponentsToContext(context.Background(), enabledComponents)

	// Test multiple enabled components
	assert.NotPanics(t, func() {
		SelectiveInfo(ctx, "search_flow", "search started")
		SelectiveInfo(ctx, "ranking_service", "ranking products")
		SelectiveInfo(ctx, "ads_integration", "this should not log")
	})
}

func TestConstants(t *testing.T) {
	// Test that constants are properly defined
	assert.Equal(t, contextKey("logging_enabled_components"), LoggingEnabledComponentsKey)
	assert.Equal(t, "Component", Component)
}

func TestConcurrentAccess(t *testing.T) {
	// Setup main logger first
	setupTestLogger()

	enabledComponents := []string{"concurrent_component"}
	ctx := AddEnabledLoggingComponentsToContext(context.Background(), enabledComponents)

	// Test concurrent access to logging functions
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() { done <- true }()

			SelectiveInfo(ctx, "concurrent_component", "concurrent message from goroutine", id)
			SelectiveDebug(ctx, "concurrent_component", "concurrent debug from goroutine", id)
			SelectiveWarn(ctx, "concurrent_component", "concurrent warn from goroutine", id)
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// If we reach here without deadlocks or panics, the test passes
	assert.True(t, true)
}
