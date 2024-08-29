package cache

import (
	"context"
	"github.com/Honeymoond24/tender-analysis/cmd/app/config"
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	client *redis.Client
	logger application.Logger
}

func NewCacheClient(address config.RedisAddress, password config.RedisPassword, logger application.Logger) Cache {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     string(address),
		Password: string(password),
		DB:       0, // use default DB
	})
	return Cache{client: redisClient, logger: logger}
}

func (c *Cache) Get(ctx context.Context, key string) (value string, err error) {
	value, err = c.client.Get(ctx, key).Result()
	if err != nil {
		c.logger.Error("Error while getting value from cache", err)
	}
	return
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (err error) {
	err = c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		c.logger.Error("Error while setting value in cache", err)
	}
	return
}
