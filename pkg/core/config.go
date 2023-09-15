package core

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"time"
)

type Config struct {
	Year int `json:"year"`
}

func LoadConfig() (*Config, error) {
	baseDir, err := GetBaseDir()
	if err != nil {
		return nil, err
	}

	configFile := path.Join(baseDir, "config.json")

	content, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return createConfig()
		}
		return nil, err
	}

	config := Config{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *Config) SaveConfig() error {
	baseDir, err := GetBaseDir()
	if err != nil {
		return err
	}

	configFile := path.Join(baseDir, "config.json")

	json, err := json.Marshal(c)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFile, json, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) String() string {
	return fmt.Sprintf("{ Year: %d, }", c.Year)
}

func createConfig() (*Config, error) {
	config := Config{Year: time.Now().Year()}

	err := config.SaveConfig()
	if err != nil {
		return nil, err
	}

	return &config, nil
}
