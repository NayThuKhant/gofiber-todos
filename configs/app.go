package configs

import (
	"os"
	"strconv"
)

const (
	ENV_DEVELOPMENT string = "development"
	ENV_PRODUCTION  string = "production"
)

type AppConfig struct {
	ENV  string
	PORT int
}

var App AppConfig

func GetEnvironment() string {
	if env := os.Getenv("ENV"); env == ENV_DEVELOPMENT || env == ENV_PRODUCTION {
		return env
	}

	panic("Invalid ENV value in .env file, ENV should be either 'development' or 'production'")
}

func IsDevelopment() bool {
	return GetEnvironment() == ENV_DEVELOPMENT
}

func IsProduction() bool {
	return GetEnvironment() == ENV_PRODUCTION
}

func InitAppConfig() {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic("Error converting PORT to int: " + err.Error())
	}

	App = AppConfig{
		ENV:  GetEnvironment(),
		PORT: port,
	}
}
