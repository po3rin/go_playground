package apilog

import (
	"context"
	"os"

	"github.com/knq/sdhook"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	logging "google.golang.org/api/logging/v2"
)

// Log - logrus struct
var Log *logrus.Logger

func init() {
	Log = logrus.New()
	Log.Level = logrus.ErrorLevel

	ctx := context.Background()
	client, err := google.DefaultClient(ctx, logging.LoggingWriteScope)
	if err != nil {
		panic(err)
	}

	resLabels := map[string]string{
		"project_id":  os.Getenv("LOG_PROJECT_ID"),
		"zone":        os.Getenv("LOG_ZONE"),
		"instance_id": "dummyinstanceid",
	}

	hook, err := sdhook.New(
		sdhook.HTTPClient(client),
		sdhook.ProjectID(os.Getenv("LOG_PROJECT_ID")),
		sdhook.LogName("test"),
		sdhook.Resource("gce_instance", resLabels),
	)
	if err != nil {
		panic(err)
	}
	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}
	Log.AddHook(hook)
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
