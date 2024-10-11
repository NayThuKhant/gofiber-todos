package configs

import "github.com/joho/godotenv"

func Init() {

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	InitAppConfig()
	InitDatbaseConfig()
}
