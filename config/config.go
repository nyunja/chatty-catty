package config

import (
	"fmt"
	"os"
)

type Config struct {
	APIKey string
}

func LoadConfig() (*Config, error) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY environment variable not set")
	}

	return &Config{APIKey: apiKey}, nil
}
