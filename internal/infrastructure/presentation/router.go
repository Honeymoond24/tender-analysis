package presentation

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/infrastructure/presentation/controllers"
	"net/http"
)

func setupRoutes() (mux *http.ServeMux) {
	mux = http.NewServeMux()

	mux.HandleFunc("GET /", controllers.RootController)
	mux.HandleFunc("GET /user/{id}", controllers.GetUser)
	mux.HandleFunc("GET /statistics", controllers.GeneralStatistics)

	return
}
