package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

var logger, _ = zap.NewDevelopment()

type Config struct {
	MacLibraryBotToken string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	return &Config{
		MacLibraryBotToken: os.Getenv("MACLIBRARY_BOT_TOKEN"),
	}, nil
}
