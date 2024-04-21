package config

import (
	"boilerplate/env"
	"boilerplate/util"
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	DBHost    string `json:"dbHost"`
	DBPort    int    `json:"dbPort"`
	Endpoints []struct {
		URL      string `json:"url"`
		AuthType string `json:"authType"`
	} `json:"endpoints"`
}

var AppConfiguration = &Config{}

func New() *Config {
	configFileName := "config.json"
	if !util.IsStringEmpty(env.AppMode) {
		configFileName = "config." + env.AppMode + ".json"

	}
	configFileS, err := os.ReadFile(filepath.Join(env.AppWd, "config", configFileName))
	if err != nil {
		slog.Error("error while loading config file", "file", configFileName, err)
	}
	json.Unmarshal(configFileS, AppConfiguration)
	return AppConfiguration
}
