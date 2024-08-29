package presentation

import (
	"github.com/Honeymoond24/tender-analysis/cmd/app/config"
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"github.com/redis/go-redis/v9"
	"net/http"
)

func Logging(next http.Handler, log application.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request: ", r.Method, " ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func NewCacheClient(address config.RedisAddress, password config.RedisPassword) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     string(address),
		Password: string(password),
		DB:       0, // use default DB
	})
	return redisClient
}
