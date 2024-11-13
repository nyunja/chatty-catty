package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey string
}

var (
	API_KEY string
	ModelName string = "gemini-pro"
	Persona string = "Pretend you are a therapy cat called Jill. Only respond like a cat therapist. The user is %s"
	DefaultIntro string = "Hi %s! I'm Jill, a therapy cat. I'll guide you through your therapy session. Remember, I'm just a helpful assistant, not a replacement for a professional"
	EndCommand string = "end"
)

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
