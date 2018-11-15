package main

import (
	"os"

	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()

	sentryDSN := os.Getenv("SENTRY_DSN")
	client, err := raven.New(sentryDSN)
	if err != nil {
		log.Fatal(err)
	}

	hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})

	if err == nil {
		log.Hooks.Add(hook)
		log.Error("WWWW")
	}
}
