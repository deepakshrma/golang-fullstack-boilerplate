package config

import (
	"log/slog"
	"os"
	"strings"
)

func getLogLevel(logLevel string) slog.Level {
	switch strings.ToLower(logLevel) {
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

func NewLogger() *slog.Logger {
	logLevelS := GetEnvD("LOG_LEVEL", "info")
	logLevel := getLogLevel(logLevelS)
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewTextHandler(os.Stdout, opts)
	return slog.New(handler)
}
