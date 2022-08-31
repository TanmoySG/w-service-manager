package config

import (
	"fmt"
	"os"
)

func LoadConfigFromFile(filepath string) (*Config, error) {
	configBytes, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("error parsing file : %s", err)
	}

	config, err := UnmarshalConfig(configBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing config JSON : %s", err)
	}

	return &config, nil
}

func LoadConfig(configBytes []byte) (*Config, error) {
	config, err := UnmarshalConfig(configBytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing config JSON : %s", err)
	}

	return &config, nil
}
