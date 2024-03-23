package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

var errLoadEnv = "Error on load env `%s`"

type Config struct {
	DbDSN     string
	DbTimeout time.Duration
}

func NewConfig(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	dbTimeout, err := strconv.Atoi(os.Getenv("DB_TIMEOUT"))
	if err != nil {
		return nil, fmt.Errorf(errLoadEnv, "db_timeout")
	}

	return &Config{
		DbDSN:     os.Getenv("DB_DSN"),
		DbTimeout: time.Second * time.Duration(dbTimeout),
	}, nil
}
