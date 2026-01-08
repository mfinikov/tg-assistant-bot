package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// BotToken is a global var to store bot token
var BotToken string

// EnvLoad is a function to load variables from .env
func EnvLoad() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	BotToken = os.Getenv("BOT_TOKEN")
}
