package config

import (
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	GoogleMapsAPIKey string
	TelegramBotToken string
	TelegramChatID string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		GoogleMapsAPIKey: os.Getenv("GOOGLE_MAPS_API_KEY"),
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		TelegramChatID: os.Getenv("TELEGRAM_CHAT_ID"),
	}

	return cfg, nil
}