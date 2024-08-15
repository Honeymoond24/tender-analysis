package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type HTTPServerPort string

func GetHTTPServerPort(log *zap.Logger) HTTPServerPort {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return HTTPServerPort(os.Getenv("API_PORT"))
}

type DatabaseDSN string

func GetDatabaseDSN(log *zap.Logger) DatabaseDSN {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return DatabaseDSN(os.Getenv("DB_DSN"))
}
