package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbDSN string
}

func NewConfig(filenames ...string) *Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		DbDSN: os.Getenv("DB_DSN"),
	}
}
