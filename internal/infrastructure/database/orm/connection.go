package orm

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/cmd/app/config"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(dsn config.DatabaseDSN, log application.Logger) *gorm.DB {
	db, err := gorm.Open(postgres.Open(string(dsn)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Info("Connected to database", zap.String("DSN", string(dsn)))
	return db
}
