package presentation

import (
	"context"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/presentation/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
	"net/http"
	"time"
)

func SetupServerHandler(apiRouter *http.ServeMux) http.Handler {
	apiPrefix := "/api/v1/"
	//apiRouter := NewServeMux()

	rootMuxRouter := http.NewServeMux()

	apiPrefixSliced := apiPrefix[:len([]rune(apiPrefix))-1] // remove the last slash: '/api/v1/' to '/api/v1'
	rootMuxRouter.Handle(apiPrefix, http.StripPrefix(apiPrefixSliced, apiRouter))
	handler := middleware.Logging(rootMuxRouter)
	return handler
}

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux, log *zap.Logger) *http.Server {
	port := "8000"
	log.Info("Starting server on port", zap.String("port", port))

	srv := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server at", zap.String("addr", srv.Addr))
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					log.Info("NewHTTPServer error", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}
