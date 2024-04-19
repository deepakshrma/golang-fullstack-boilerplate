package env

import (
	"boilerplate/util"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

var APP_MODE = ""

func GetEnvD(key string, def string) string {
	env := os.Getenv(key)
	if util.IsStringEmpty(env) {
		return def
	}
	return env
}

func LoadEnvs() {
	env := GetEnvD("APP_MODE", "")
	APP_MODE = env
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
		fmt.Println(params)
		key := strings.TrimSpace(params[0])
		if key == "" {
			continue
		}
		val := strings.TrimSpace(params[1])
		os.Setenv(key, val)
	}
}
