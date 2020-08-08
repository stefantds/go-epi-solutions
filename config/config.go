package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	TestDataFolder string `yaml:"testDataFolder"`
}

// Parse parses the config file and returns the configuration object
func Parse() (*Config, error) {
	pwd, err := os.Getwd()

	f, err := os.Open(filepath.Join(pwd, "../../config.yml"))
	if err != nil {
		return nil, fmt.Errorf("can't find config file: %w", err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)

	var cfg Config
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("can't parse config file: %w", err)
	}

	return &cfg, nil
}