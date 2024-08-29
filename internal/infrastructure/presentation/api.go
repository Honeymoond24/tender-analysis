package presentation

import (
	"context"
	"github.com/Honeymoond24/tender-analysis/cmd/app/config"
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"go.uber.org/fx"
	"net"
	"net/http"
	"time"
)

//func SetupServerHandler(apiRouter *http.ServeMux) http.Handler {
//	apiPrefix := "/api/v1/"
//	//apiRouter := NewServeMux()
//
//	rootMuxRouter := http.NewServeMux()
//
//	apiPrefixSliced := apiPrefix[:len([]rune(apiPrefix))-1] // remove the last slash: '/api/v1/' to '/api/v1'
//	rootMuxRouter.Handle(apiPrefix, http.StripPrefix(apiPrefixSliced, apiRouter))
//	handler := middleware.Logging(rootMuxRouter)
//	return handler
//}

func NewHTTPServer(
	lc fx.Lifecycle,
	mux *http.ServeMux,
	log application.Logger,
	port config.HTTPServerPort,
) *http.Server {
	log.Info("Starting server on port", "port", string(port))

	srv := &http.Server{
		Addr:           string(":" + port),
		Handler:        mux,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			log.Info("Starting HTTP server at", "addr", srv.Addr)
			go func() {
				err := srv.Serve(ln)
				if err != nil {
					log.Info("NewHTTPServer error", err)
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
