package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DbURL string `json:"db_url"`
	User  string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	path := filepath.Join(home_dir, configFileName)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("%s file not exists", path)
	}

	return fmt.Sprintf("%s/%s", home_dir, configFileName), nil
}

func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, nil
	}
	c, err := os.ReadFile(path)
	if err != nil {
		return Config{}, nil
	}
	config := Config{}
	if err := json.Unmarshal(c, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (config *Config) SetUser(name string) error {
	// config, err := Read()
	// if err != nil {
	// 	return err
	// }
	config.User = name
	c, err := json.Marshal(config)
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	os.WriteFile(path, c, 0777)
	return nil
}
