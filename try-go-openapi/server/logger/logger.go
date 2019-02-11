package logger

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

// Log - logrus struct
var Log *logrus.Logger

func init() {
	Log = logrus.New()
	switch os.Getenv("LOG_LEVEL") {
	case "DEBUG":
		Log.Level = logrus.DebugLevel
	case "INFO":
		Log.Level = logrus.InfoLevel
	case "WARN":
		Log.Level = logrus.WarnLevel
	case "ERROR":
		Log.Level = logrus.ErrorLevel
	default:
		Log.Level = logrus.ErrorLevel
	}

	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}

	sentryDSN := os.Getenv("SENTRY_DSN")
	client, err := raven.New(sentryDSN)
	if err != nil {
		Fatal(err)
	}
	hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	})

	if err == nil {
		Log.Hooks.Add(hook)
	}
}

// Debug - shorthand debug
func Debug(args ...interface{}) {
	Log.Debug(args...)
}

// Info - shorthand info
func Info(args ...interface{}) {
	Log.Info(args...)
}

// Warn - shorthand warn
func Warn(args ...interface{}) {
	Log.Warn(args...)
}

// Error - shorthand error
func Error(args ...interface{}) {
	Log.Error(args...)
}

// Fatal - shorthand fatal
func Fatal(args ...interface{}) {
	Log.Fatal(args...)
}

// Panic - shorthand panic
func Panic(args ...interface{}) {
	Log.Panic(args...)
}
