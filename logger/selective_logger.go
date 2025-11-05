package logger

import (
	"context"
	"strings"

	"github.com/sirupsen/logrus"
)

type contextKey string

const (
	LoggingEnabledComponentsKey contextKey = "logging_enabled_components"
	Component                   string     = "Component"
)

func AddEnabledLoggingComponentsToContext(ctx context.Context, enabledComponents []string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, LoggingEnabledComponentsKey, enabledComponents)
}

func isComponentEnabled(ctx context.Context, componentName string) bool {
	if ctx == nil {
		return false
	}
	enabledComponents, ok := ctx.Value(LoggingEnabledComponentsKey).([]string)
	if !ok {
		return false
	}

	for _, component := range enabledComponents {
		if strings.EqualFold(component, componentName) {
			return true
		}
	}

	return false
}

func getComponentLogger(ctx context.Context, componentName string) *logrus.Entry {
	return GetLoggerFromContext(ctx).WithField(Component, componentName)
}

func SelectiveDebug(ctx context.Context, componentName string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Debug(args...)
	}
}

func SelectiveDebugf(ctx context.Context, componentName string, format string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Debugf(format, args...)
	}
}

func SelectiveInfo(ctx context.Context, componentName string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Info(args...)
	}
}

func SelectiveInfof(ctx context.Context, componentName string, format string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Infof(format, args...)
	}
}

func SelectiveWarn(ctx context.Context, componentName string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Warn(args...)
	}
}

func SelectiveWarnf(ctx context.Context, componentName string, format string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Warnf(format, args...)
	}
}

func SelectiveError(ctx context.Context, componentName string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Error(args...)
	}
}

func SelectiveErrorf(ctx context.Context, componentName string, format string, args ...interface{}) {
	if isComponentEnabled(ctx, componentName) {
		getComponentLogger(ctx, componentName).Errorf(format, args...)
	}
}
