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
	Log.Formatter = &logrus.TextFormatter{FullTimestamp: true, DisableColors: true}

	if os.Getenv("EXECUTION_ENVIRONMENT") == "gce" {
		ctx := context.Background()

		client, err := google.DefaultClient(ctx, logging.LoggingWriteScope)
		if err != nil {
			panic(err)
		}

		resLabels := map[string]string{
			"project_id": os.Getenv("GCP_PROJECT_ID"),
			"zone":       "asia-northeast1-b",
		}

		hook, err := sdhook.New(
			sdhook.HTTPClient(client),
			sdhook.ProjectID(os.Getenv("GCP_PROJECT_ID")),
			sdhook.LogName("test"),
			sdhook.Resource("gce_instance", resLabels),
		)
		if err != nil {
			panic(err)
		}
		Log.AddHook(hook)
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
