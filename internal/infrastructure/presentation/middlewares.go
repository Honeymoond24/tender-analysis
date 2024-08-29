package presentation

import (
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"net/http"
)

func Logging(next http.Handler, log application.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("Request: ", r.Method, " ", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
