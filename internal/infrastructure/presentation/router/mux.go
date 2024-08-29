package router

import (
	"net/http"
)

type Route interface {
	http.Handler
	Pattern() string
}

func NewServeMux(routes []Route) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	for _, route := range routes {
		//handler := Chain(route.Middlewares(), route)
		mux.Handle(route.Pattern(), route)
	}
	return
}
