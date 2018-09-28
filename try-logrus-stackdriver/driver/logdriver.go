package apilog

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log - logrus struct
var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Level = logrus.ErrorLevel
	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}
	if os.Getenv("EXECUTION_ENVIRONMENT") == "gce" {
		Log.Out = os.Stdout
		return
	}
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	Log.Out = logFile
}

// Debug - shorthand Log.Debug
func Debug(msg string) {
	Log.Debug(msg)
}

// Info - shorthand Log.Info
func Info(msg string) {
	Log.Info(msg)
}

// Warn - shorthand Log.Warn
func Warn(msg string) {
	Log.Warn(msg)
}

// Error - shorthand Log.Error
func Error(msg string) {
	Log.Error(msg)
}
