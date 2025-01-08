package config

import (
	"log/slog"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type AppConfig struct {
	Port     string `yaml:"port" default:"8181"`
	APIKey   string `yaml:"api_key"`
	Login    string `yaml:"login"`
	Password string `yaml:"password"`
}

type BackupConfig []struct {
	DBName  string
	DBType  string
	path    string
	shedule string
}

func ReadConfig(configfile string) (*AppConfig, error) {
	if configfile == "" {
		configfile = "app.yml"
	}

	data, err := os.ReadFile(configfile)
	if err != nil {
		slog.Error("Error reading config file: ", err)
	}

	config := &AppConfig{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		slog.Error("Error unmarshalling config file: ", err)
	}
	return config, nil
}
