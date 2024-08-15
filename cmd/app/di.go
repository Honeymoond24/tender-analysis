package main

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/cmd/app/config"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/database/orm"
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
			config.GetHTTPServerPort,                                        // HTTPServerPort
			config.GetDatabaseDSN,                                           // DatabaseDSN
			orm.Connection,                                                  // *gorm.DB
			orm.NewStatisticsRepository,                                     // application.Statistics
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	}
}
