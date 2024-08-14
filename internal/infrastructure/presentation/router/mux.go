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
		mux.Handle(route.Pattern(), route)
	}
	return
}
