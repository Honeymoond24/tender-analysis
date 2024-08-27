package main

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/cmd/app/config"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/database/adapter"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/logs"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/presentation"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/presentation/router"
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

func AsRouteWithLogging(f any) any {
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
			AsRouteWithLogging(router.NewPingHandler),                       // *TestResponseTimeHandler
			//presentation.SetupServerHandler,                                 // http.Handler
			logs.NewLogger,                  // *Logger
			zap.NewProduction,               // *zap.Logger for fx
			config.GetHTTPServerPort,        // HTTPServerPort
			config.GetDatabaseDSN,           // DatabaseDSN
			adapter.NewPG,                   // *adapter.DBPool
			adapter.NewStatisticsRepository, // application.Statistics
			presentation.NewCacheClient,     // *redis.Client
			config.GetRedisAddress,          // RedisAddress
			config.GetRedisPassword,         // RedisPassword
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	}
}
