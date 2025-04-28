package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:db_url`
	CurrentUserName string `json:current_user_name`
}

const configFileName = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get home directory - %v", err)
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

// This is a internal function so not exported
func write(config Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("unable to open the file - %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return fmt.Errorf("error writing file - %v", err)
	}
	return nil
}

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	var config Config

	if err != nil {
		return config, fmt.Errorf("error retrieving home directory - %v", err)
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return config, fmt.Errorf("unable to read the config file %s - %v", fullPath, err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("incorrect config format -  %v", err)
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName
	return write(*c)
}
