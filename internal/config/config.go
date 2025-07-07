package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Read decodes the JSON config file into a Config struct.
func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	config := Config{}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&config); err != nil {
		return Config{}, err

	}

	return config, nil
}

// SetUser updates the current user's name and saves the configuration.
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}

// write encodes the given Config struct to the JSON config file.
func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	// Sets indentation for pretty-printing the JSON
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(cfg); err != nil {
		return err
	}

	return nil
}

// getConfigFilePath builds and returns the absolute path to the config file.
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filePath := filepath.Join(home, configFileName)
	return filePath, nil
}
