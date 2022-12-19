package logging

import (
	"api/config"

	"github.com/sirupsen/logrus"
)

var logger = initLogger()

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logLevel, err := logrus.ParseLevel(config.Env.LogLevel)
	if err != nil {
		logger.SetLevel(logrus.DebugLevel)
		logger.Warn("Error parsing log level from environment")
	}
	logger.SetLevel(logLevel)
	return logger
}

func Trace(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Trace(args...)
}

func Debug(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Debug(args...)
}

func Info(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Info(args...)
}

func Warn(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Warn(args...)
}

func Error(fields logrus.Fields, args ...interface{}) {
	logger.WithFields(fields).Error(args...)
}

func Fatal(fields logrus.Fields, args ...interface{}) {
	// Calls os.Exit(1) after logging
	logger.WithFields(fields).Fatal(args...)
}

func Panic(args ...interface{}) {
	// Calls panic() after logging
	logger.Panic(args...)
}
