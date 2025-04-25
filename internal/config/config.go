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

func Read() (Config, error) {
	home_dir, err := os.UserHomeDir()
	var config Config

	if err != nil {
		return config, fmt.Errorf("Error retriving home directory - %v", err)
	}

	data, err := os.ReadFile(home_dir + string(filepath.Separator) + ".gatorconfig.json")
	if err != nil {
		return config, fmt.Errorf("Unable to read the config file - %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("Incorrect config format -  %v", err)
	}

	return config, nil
}

func (c *Config) SetUser(userName string) error {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return fmt.Errorf("Error retriving home directory - %v", err)
	}
	filePath := filepath.Join(homeDir, ".gatorconfig.json")
	c.Current_user_name = userName

	newData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("Error encoding JSON - %v", err)
	}

	err = os.WriteFile(filePath, newData, 0664)
	if err != nil {
		return fmt.Errorf("Error writing file - %v", err)
	}
	return nil
}
