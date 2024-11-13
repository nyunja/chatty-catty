package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err!= nil {
        return nil, fmt.Errorf("error loading .env file: %v", err)
    }
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("API_KEY environment variable not set")
	}

	return &Config{APIKey: apiKey}, nil
}
