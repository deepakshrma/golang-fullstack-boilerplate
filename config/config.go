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
		Url      string `json:"url"`
		AuthType string `json:"authType"`
	} `json:"endpoints"`
}

var AppConfiguration = &Config{}

func New() *Config {
	configFileName := "config.json"
	if !util.IsStringEmpty(env.APP_MODE) {
		configFileName = "config." + env.APP_MODE + ".json"

	}
	configFileS, err := os.ReadFile(filepath.Join(env.APP_WD, "config", configFileName))
	if err != nil {
		slog.Error("error while loading config file", "file", configFileName, err)
	}
	json.Unmarshal(configFileS, AppConfiguration)
	return AppConfiguration
}
