package presentation

import (
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type Route interface {
	http.Handler
	Pattern() string
}

//func NewServeMux(rootRoute, statsRoute Route) (mux *http.ServeMux) {
//	mux = http.NewServeMux()
//
//	mux.Handle(rootRoute.Pattern(), rootRoute)
//	mux.HandleFunc("GET /user/{id}", GetUser)
//	mux.Handle(statsRoute.Pattern(), statsRoute)
//
//	return
//}

func NewServeMux(routes []Route) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	for _, route := range routes {
		mux.Handle(route.Pattern(), route)
	}
	return
}

type RootHandler struct {
	log *zap.Logger
}

func NewRootHandler(log *zap.Logger) *RootHandler {
	return &RootHandler{log: log}
}

func (h *RootHandler) Pattern() string {
	return "/"
}

func (h *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RootHandler.ServeHTTP", r.RequestURI)
	h.log.Info("RootHandler.ServeHTTP", zap.String("RequestURI", r.RequestURI))

	_, err := w.Write([]byte(`{"message": "RootHandler working..."}`))
	if err != nil {
		h.log.Warn("RootHandler.ServeHTTP error", zap.Error(err))
		return
	}
}

type StatisticsHandler struct {
	log *zap.Logger
}

func NewStatisticsHandler(log *zap.Logger) *StatisticsHandler {
	return &StatisticsHandler{log: log}
}

func (h *StatisticsHandler) Pattern() string {
	return "/statistics"
}

func (h *StatisticsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Error("StatisticsHandler.ServeHTTP error", zap.Error(err))
	}
	if _, err := fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI, "GetUser user id - ", r.PathValue("id"))

	_, err := w.Write([]byte(`{"message": "working..."}`))
	if err != nil {
		return
	}
}
