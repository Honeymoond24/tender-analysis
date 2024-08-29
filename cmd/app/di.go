package main

import (
	"github.com/Honeymoond24/tender-analysis/cmd/app/config"
	"github.com/Honeymoond24/tender-analysis/internal/infrastructure/database/adapter"
	"github.com/Honeymoond24/tender-analysis/internal/infrastructure/logs"
	"github.com/Honeymoond24/tender-analysis/internal/infrastructure/presentation"
	"github.com/Honeymoond24/tender-analysis/internal/infrastructure/presentation/router"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(router.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func GetFxOptions() []fx.Option {
	return []fx.Option{
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			presentation.NewHTTPServer,                                      // *http.Server
			fx.Annotate(router.NewServeMux, fx.ParamTags(`group:"routes"`)), // *http.ServeMux
			AsRoute(router.NewStatisticsHandler),                            // *StatisticsHandler
			AsRoute(router.NewPersonalStatisticsHandler),                    // *PersonalStatisticsHandler
			presentation.SetupServerHandler,                                 // http.Handler
			logs.NewLogger,                                                  // *Logger
			zap.NewProduction,                                               // *zap.Logger for fx
			config.GetHTTPServerPort,                                        // HTTPServerPort
			config.GetDatabaseDSN,                                           // DatabaseDSN
			adapter.NewPG,                                                   // *adapter.DBPool
			adapter.NewStatisticsRepository,                                 // application.Statistics
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	}
}
