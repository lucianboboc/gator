package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(fileBytes, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	err := write(c)
	if err != nil {
		return err
	}
	return nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := homeDir + "/" + configFileName
	return filePath, nil
}

func write(cfg *Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, bytes, 0644)
}
