package router

import (
	"fmt"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application"
	"net/http"
)

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
	statistics := application.GetPersonalStatistics(userId)
	if _, err := fmt.Fprint(w, statistics); err != nil {
		h.log.Error("Failed to write response", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
