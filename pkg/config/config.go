package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"webapp/pkg/env"
	"webapp/pkg/helpers/str"
)

type AppConfig struct {
	Port      int    `json:"port"`
	DBHost    string `json:"dbHost"`
	DBPort    int    `json:"dbPort"`
	Endpoints []struct {
		URL      string `json:"url"`
		AuthType string `json:"authType"`
	} `json:"endpoints"`
}

func NewAppConfig() (*AppConfig, error) {
	configFileName := "config.json"
	if !str.IsStringEmpty(env.AppMode) {
		configFileName = "config." + env.AppMode + ".json"
	}
	configFileS, err := os.ReadFile(filepath.Join(env.AppWd, "config", configFileName))
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	config := &AppConfig{}
	err = json.Unmarshal(configFileS, config)
	if config.Port == 0 {
		config.Port = 8080
	}
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	return config, nil
}
