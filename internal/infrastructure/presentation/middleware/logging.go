package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		defer log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
		next.ServeHTTP(w, req)
	})
}
