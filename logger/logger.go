package logger

import (
	"base_app/config"
	"base_app/constants/api_constants"
	"context"
	"os"
	"time"

	"github.com/evalphobia/logrus_sentry"
	"github.com/google/uuid"
	joonix "github.com/joonix/log"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

var Log *Logger

type LoggerError struct {
	Error error
}

func panicIfError(err error) {
	if err != nil {
		panic(LoggerError{err})
	}
}

func SetupLogger() {
	level, err := logrus.ParseLevel(config.LogLevel())
	panicIfError(err)

	logrusVar := &logrus.Logger{
		Out:       os.Stdout,
		Formatter: joonix.NewFormatter(),
		Hooks:     make(logrus.LevelHooks),
		Level:     level,
	}

	Log = &Logger{logrusVar}

	if config.EnableSentry() {
		sentryLevels := []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
		}
		sentryHook, err := logrus_sentry.NewSentryHook(config.SentryDsn(), sentryLevels)
		sentryHook.Timeout = 1 * time.Second
		panicIfError(err)

		Log.Hooks.Add(sentryHook)
	}
}

func BuildContext(context string) logrus.Fields {
	return logrus.Fields{
		"context":    context,
		"request_id": uuid.NewString(),
	}
}

func GetStringFromContext(ctx context.Context, key interface{}) string {
	if val, ok := ctx.Value(key).(string); ok {
		return val
	}
	return ""
}

func GetLoggerFromContext(ctx context.Context) *logrus.Entry {
	if ctx == nil {
		return Log.WithFields(logrus.Fields{})
	}
	fields := logrus.Fields{
		api_constants.RequestId:     GetStringFromContext(ctx, api_constants.RequestId),
		api_constants.UserId:        GetStringFromContext(ctx, api_constants.UserId),
		api_constants.UserSessionId: GetStringFromContext(ctx, api_constants.UserSessionId),
		api_constants.DeviceId:      GetStringFromContext(ctx, api_constants.DeviceId),
	}
	return Log.WithFields(fields)
}

func Debug(ctx context.Context, args ...interface{}) {
	GetLoggerFromContext(ctx).Debug(args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	GetLoggerFromContext(ctx).Debugf(format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	GetLoggerFromContext(ctx).Info(args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	GetLoggerFromContext(ctx).Infof(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	GetLoggerFromContext(ctx).Error(args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	GetLoggerFromContext(ctx).Errorf(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	GetLoggerFromContext(ctx).Warn(args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	GetLoggerFromContext(ctx).Warnf(format, args...)
}
