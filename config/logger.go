package config

import (
	"boilerplate/env"
	"log/slog"
	"os"
)

func getLogLevel(logLevel string) slog.Level {
	switch logLevel {
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

var Logger *slog.Logger

func InitLogger() {
	logLevelS := env.GetEnvD("LOG_LEVEL", "info")
	logLevel := getLogLevel(logLevelS)
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	Logger = slog.New(handler)
}
