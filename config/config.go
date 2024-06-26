package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerUrl string
	DbDSN     string
	DbTimeout time.Duration
}

func NewConfig(filenames ...string) (*Config, error) {
	err := godotenv.Load(filenames...)
	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	dbTimeout, err := strEnvToInt("DB_TIMEOUT")
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerUrl: os.Getenv("SERVER_URL"),
		DbDSN:     os.Getenv("DB_DSN"),
		DbTimeout: time.Second * time.Duration(dbTimeout),
	}, nil
}

func strEnvToInt(envName string) (int, error) {
	env, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		return 0, fmt.Errorf("Error on load env `%s`", env)
	}
	return env, nil
}
