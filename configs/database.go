package configs

import (
	"os"
	"strconv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

var Database DatabaseConfig

func InitDatbaseConfig() {
	databasePort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Error converting DB_PORT to int: " + err.Error())
	}

	Database = DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     databasePort,
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
}
