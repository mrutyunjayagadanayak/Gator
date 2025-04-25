package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DB_Url            string `json:db_url`
	Current_user_name string `json:current_user_name`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	fullPath, err := getConfigFilePath()
	var config Config

	if err != nil {
		return config, fmt.Errorf("Error retriving home directory - %v", err)
	}

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return config, fmt.Errorf("Unable to read the config file - %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("Incorrect config format -  %v", err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Unable to get Home directory - %v", err)
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}

func write(config Config) error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return fmt.Errorf("Unable to open the file - %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(config)
	if err != nil {
		return fmt.Errorf("Error writing file - %v", err)
	}
	return nil
}

func (c *Config) SetUser(userName string) error {
	c.Current_user_name = userName
	return write(*c)
}
