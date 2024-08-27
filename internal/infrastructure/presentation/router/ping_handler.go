package router

import (
	"fmt"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type PingHandler struct {
	log         application.Logger
	cacheClient *redis.Client
}

func NewPingHandler(log application.Logger, cacheClient *redis.Client) *PingHandler {
	pingHandler := &PingHandler{log: log, cacheClient: cacheClient}
	return pingHandler
}

func (h *PingHandler) Pattern() string {
	return "/test"
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	value, err := h.cacheClient.Get(r.Context(), h.Pattern()).Result()
	if err == nil && value != "" {
		h.log.Info(fmt.Sprintf("Serving from cache %v %v", h.Pattern(), value))
		_, _ = fmt.Fprint(w, value)
		return
	}

	responseBody := application.TestResponseTime()
	fmt.Println("TestResponseTime", responseBody, h.Pattern())

	if _, err := fmt.Fprint(w, responseBody); err != nil {
		h.log.Error("Failed to write response", r.RequestURI, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	h.cacheClient.Set(r.Context(), h.Pattern(), responseBody, 5*time.Second)
}
