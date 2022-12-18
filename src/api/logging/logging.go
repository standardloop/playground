package logging

import (
	"api/config"

	"github.com/sirupsen/logrus"
)

var logger = initLogger()

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(config.Env.LogLevel)
	if err != nil {
		logger.SetLevel(logrus.DebugLevel)
		logger.Warn("Error parsing log level from environment")
	}
	logger.SetLevel(logLevel)
	return logger
}

func Trace(args ...interface{}) {
	logger.Trace(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	// Calls os.Exit(1) after logging
	logger.Fatal(args...)
}

func Panic(args ...interface{}) {
	// Calls panic() after logging
	logger.Panic(args...)
}
