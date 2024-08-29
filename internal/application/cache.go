package application

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (value string, err error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error)
}
