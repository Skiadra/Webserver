package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUri      string
	ServerPort string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		DBUri:      os.Getenv("DB_URI"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}
