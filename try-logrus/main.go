package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	l.Info("ii")
	l.Error("ee")
}

// l - logrus struct
var l *logrus.Logger

func init() {
	l = logrus.New()
	l.Level = logrus.ErrorLevel

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	l.Formatter = &logrus.JSONFormatter{}
	l.Out = logFile
}

// Debug - shorthand l.Debug
func Debug(msg string) {
	l.Debug(msg)
}

// Info - shorthand l.Info
func Info(msg string) {
	l.Info(msg)
}

// Warn - shorthand l.Warn
func Warn(msg string) {
	l.Warn(msg)
}

// Error - shorthand l.Error
func Error(msg string) {
	l.Error(msg)
}
