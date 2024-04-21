package env

import (
	"boilerplate/util"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var AppMode = ""
var AppWd = "./"

func GetEnvD(key string, def string) string {
	env := os.Getenv(key)
	if util.IsStringEmpty(env) {
		return def
	}
	return env
}

func LoadEnvs() {
	env := GetEnvD("APP_MODE", "")
	AppMode = env
	AppWd, _ = os.Getwd()
	envPath := "./env/application.env"
	if !util.IsStringEmpty(env) {
		envPath = fmt.Sprintf("./env/application.%s.env", env)
	}
	absPath, err := filepath.Abs(envPath)
	if err != nil {
		slog.Error("unable to find " + envPath)
	}

	envFile, err := os.ReadFile(absPath)
	if err != nil {
		slog.Error("unable to load " + absPath)
	}
	for _, line := range strings.Split(strings.TrimSpace(string(envFile)), "\n") {
		params := strings.Split(line, "=")
		key := strings.TrimSpace(params[0])
		if key == "" {
			continue
		}
		val := strings.TrimSpace(params[1])
		err := os.Setenv(key, val)
		if err != nil {
			slog.Warn("unable to set env ", "key", key)
		}
	}
}
