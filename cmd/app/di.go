package main

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/cmd/app/config"
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
			presentation.NewHTTPServer, // http server
			fx.Annotate(
				router.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			), // http serve mux with routes
			AsRoute(router.NewStatisticsHandler),
			AsRoute(router.NewPersonalStatisticsHandler),
			presentation.SetupServerHandler,
			zap.NewProduction, // logger
			config.GetHTTPServerPort,
			config.GetDatabaseDSN,                                           // DatabaseDSN
			orm.Connection,                                                  // *gorm.DB
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	}
}
