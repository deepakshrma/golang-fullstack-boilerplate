package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"webapp/pkg/helpers/str"
)

var AppMode = ""
var AppWd = "./"

func GetEnvD(key string, def string) string {
	env := os.Getenv(key)
	if str.IsStringEmpty(env) {
		return def
	}
	return env
}

func LoadEnvs() {
	env := GetEnvD("APP_MODE", "")
	AppMode = env
	AppWd, _ = os.Getwd()
	envPath := "./.env"
	if !str.IsStringEmpty(env) {
		envPath = fmt.Sprintf("./.env.%s", env)
	}
	absPath, err := filepath.Abs(envPath)
	if err != nil {
		slog.Error("unable to find " + envPath)
		return
	}

	envFile, err := os.ReadFile(absPath)
	if err != nil {
		slog.Error("unable to load " + absPath)
		return
	}
	for _, line := range strings.Split(strings.TrimSpace(string(envFile)), "\n") {
		if strings.HasPrefix(line, "#") {
			continue
		}
		params := strings.Split(line, "=")
		key := strings.TrimSpace(params[0])
		if key == "" {
			continue
		}
		val := strings.TrimSpace(params[1])
		if err := os.Setenv(key, val); err != nil {
			slog.Warn("unable to set env ", "key", key)
		}
	}
}
