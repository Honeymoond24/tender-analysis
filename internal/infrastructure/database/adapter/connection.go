package adapter

import (
	"context"
	"github.com/Honeymoond24/tender-analysis/cmd/app/config"
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type DBPool struct {
	conn *pgxpool.Pool
}

var (
	pgInstance *DBPool
	pgOnce     sync.Once
)

func NewPG(dsn config.DatabaseDSN, log application.Logger) (*DBPool, error) {
	pgOnce.Do(func() {
		db, err := pgxpool.New(context.Background(), string(dsn))
		if err != nil {
			log.Fatal("Failed to connect to database", err)
		}
		pgInstance = &DBPool{conn: db}
	})

	return pgInstance, nil
}

func (pg DBPool) Ping(ctx context.Context) error {
	return pg.conn.Ping(ctx)
}

func (pg DBPool) Close() {
	pg.conn.Close()
}
