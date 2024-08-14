package config

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
)

type HTTPServerPort string

func (p HTTPServerPort) String() string {
	return string(p)
}

func GetHTTPServerPort(log *zap.Logger) HTTPServerPort {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return HTTPServerPort(os.Getenv("API_PORT"))
}
