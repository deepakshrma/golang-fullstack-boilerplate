package config

import (
	"boilerplate/env"
	"boilerplate/util"
	"encoding/json"
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
		Logger.Error("error while loading config file", "file", configFileName, err)
	}
	err = json.Unmarshal(configFileS, AppConfiguration)
	if err != nil {
		Logger.Error("error while parse config file", "file", configFileName, err)
	}
	return AppConfiguration
}
