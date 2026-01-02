package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"chat-app/internal/domain"
)

func GetDBConfig() domain.DBConfig {
	if err := godotenv.Load("config.env"); err != nil {
		log.Fatalf("error load config DB: %v", err)
	}

	return domain.DBConfig{
		User: os.Getenv("POSTGRES_USER"),
		Pass: os.Getenv("POSTGRES_PASSWORD"),
		DB:   os.Getenv("POSTGRES_DB"),
		Host: os.Getenv("POSTGRES_HOST"),
		Port: os.Getenv("POSTGRES_PORT"),
	}
}

// TODO: add func GetAddr()
