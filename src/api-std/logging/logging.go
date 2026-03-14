package logging

import (
	"api-std/config"
	"log/slog"
	"os"
)

func Init() {
	var logLevel slog.LevelVar

	switch config.Env.LogLevel {
	case "ERROR":
		logLevel.Set(slog.LevelError)
	case "WARN":
		logLevel.Set(slog.LevelWarn)
	case "DEBUG":
		logLevel.Set(slog.LevelDebug)
	case "INFO":
		logLevel.Set(slog.LevelInfo)
	default:
		logLevel.Set(slog.LevelInfo)
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: &logLevel,
	}))
	slog.SetDefault(logger)
}
