package main

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/presentation"
	"go.uber.org/fx"
	"net/http"
)

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
		fx.Provide(presentation.NewHTTPServer),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
