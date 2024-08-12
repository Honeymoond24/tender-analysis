package main

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/presentation"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"net/http"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(presentation.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}

func main() {
	//err := godotenv.Load()
	//port, err := strconv.Atoi(os.Getenv("API_PORT"))
	//apiPrefix := os.Getenv("API_PREFIX")
	//log.Println("Read environment variables")

	//err = presentation.StartServer(port, apiPrefix)
	//if err != nil {
	//	log.Fatal(err)
	//}

	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			presentation.NewHTTPServer, // http server
			fx.Annotate(
				presentation.NewServeMux,
				fx.ParamTags(`group:"routes"`),
			), // http serve mux with routes
			AsRoute(presentation.NewRootHandler),
			AsRoute(presentation.NewStatisticsHandler),
			presentation.SetupServerHandler,
			zap.NewProduction, // logger
		),
		fx.Invoke(
			func(*http.Server) {},
		),
	).Run()
}
