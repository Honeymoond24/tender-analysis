package router

import (
	"encoding/json"
	"fmt"
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"github.com/Honeymoond24/tender-analysis/internal/application/use_cases"
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
	statistics := use_cases.GetGeneralStatistics(h.repository)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(statistics)

	h.log.Info("Statistics request", r.RequestURI)
	if err != nil {
		h.log.Error("Failed to write response", err)
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
		h.log.Error("Failed to write response", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
