package config

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application"
	"github.com/joho/godotenv"
	"os"
)

type HTTPServerPort string

func GetHTTPServerPort(log application.Logger) HTTPServerPort {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return HTTPServerPort(os.Getenv("API_PORT"))
}

type DatabaseDSN string

func GetDatabaseDSN(log application.Logger) DatabaseDSN {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return DatabaseDSN(os.Getenv("DB_DSN"))
}
