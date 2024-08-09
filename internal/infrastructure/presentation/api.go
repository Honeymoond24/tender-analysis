package presentation

import (
	"context"
	"fmt"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/presentation/middleware"
	"go.uber.org/fx"
	"log"
	"net"
	"net/http"
	"time"
)

func setupHandler(apiPrefix string) http.Handler {
	apiRouter := setupRoutes()

	rootMuxRouter := http.NewServeMux()

	apiPrefixSliced := apiPrefix[:len([]rune(apiPrefix))-1] // remove the last slash: '/api/v1/' to '/api/v1'
	rootMuxRouter.Handle(apiPrefix, http.StripPrefix(apiPrefixSliced, apiRouter))

	handler := middleware.Logging(rootMuxRouter)
	return handler
}

func NewHTTPServer(lc fx.Lifecycle) *http.Server {
	port := "8000"
	apiPrefix := "/api/v1/"
	log.Println("Starting server on port", port)
	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        setupHandler(apiPrefix),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//log.Fatal(s.ListenAndServe())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			//go func() {
			//	err := srv.Serve(ln)
			//	if err != nil {
			//		log.Println(err)
			//	}
			//}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
