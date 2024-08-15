package router

import (
	"fmt"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application/use_cases"
	"go.uber.org/zap"
	"net/http"
)

// r.PathValue("id")

type StatisticsHandler struct {
	log        application.Logger
	repository application.StatisticsRepository
}

func NewStatisticsHandler(log application.Logger, repository application.StatisticsRepository) *StatisticsHandler {
	return &StatisticsHandler{log: log, repository: repository}
}
func (h *StatisticsHandler) Pattern() string {
	return "/statistics"
}

func (h *StatisticsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	statistics := use_cases.GetGeneralStatistics(&h.repository)
	_, err := fmt.Fprint(w, statistics)
	if err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

type PersonalStatisticsHandler struct {
	log application.Logger
}

func NewPersonalStatisticsHandler(log application.Logger) *PersonalStatisticsHandler {
	return &PersonalStatisticsHandler{log: log}
}
func (h *PersonalStatisticsHandler) Pattern() string {
	return "/statistics/{id}"
}

func (h *PersonalStatisticsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")
	statistics := use_cases.GetPersonalStatistics(userId)
	if _, err := fmt.Fprint(w, statistics); err != nil {
		h.log.Error("Failed to write response", zap.Error(err))
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
