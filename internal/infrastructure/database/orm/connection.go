package orm

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

func Connection(dsn string) *gorm.DB {
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	dsnWithFolder := filepath.Join("data", dsn)
	fmt.Println("DSN:", dsnWithFolder)
	db, err := gorm.Open(sqlite.Open(dsnWithFolder), &gorm.Config{})
	//db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
