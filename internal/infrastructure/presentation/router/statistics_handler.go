package router

import (
	"encoding/json"
	"fmt"
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"github.com/Honeymoond24/tender-analysis/internal/infrastructure/cache"
	"net/http"
	"strconv"
	"time"
)

type StatisticsHandler struct {
	log        application.Logger
	repository application.StatisticsRepository
	cache      cache.Cache
}

func NewStatisticsHandler(
	log application.Logger,
	repository application.StatisticsRepository,
	cache cache.Cache,
) *StatisticsHandler {
	return &StatisticsHandler{
		log:        log,
		repository: repository,
		cache:      cache,
	}
}
func (h *StatisticsHandler) Pattern() string {
	return "/statistics"
}

func (h *StatisticsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	params := application.Params{}
	var err error
	if showSum := queryParams.Get("show_sum"); showSum != "" {
		if showSum == "True" {
			params.ShowSum = true
		} else if showSum == "False" {
			params.ShowSum = false
		} else {
			http.Error(w, "show_sum must be a boolean", http.StatusUnprocessableEntity)
			return
		}
	}
	if sumRangeFrom := queryParams.Get("sum_range_from"); sumRangeFrom != "" {
		if params.SumRangeFrom, err = strconv.Atoi(sumRangeFrom); err != nil {
			http.Error(w, "sum_range_from must be an integer", http.StatusUnprocessableEntity)
			return
		}
	}
	if sumRangeTo := queryParams.Get("sum_range_to"); sumRangeTo != "" {
		if params.SumRangeTo, err = strconv.Atoi(sumRangeTo); err != nil {
			http.Error(w, "sum_range_to must be an integer", http.StatusUnprocessableEntity)
			return
		}
	}
	params.CategoryCode = queryParams.Get("category")
	params.KeyWords = queryParams["key_word"]

	fmt.Println(fmt.Sprintf("1 %#v", params))
	fmt.Println(fmt.Sprintf("2 %v", params))

	cachedResponse, err := h.cache.Get(r.Context(), h.Pattern()+fmt.Sprintf("%#v", params))
	if err == nil && cachedResponse != "" {
		h.log.Info(fmt.Sprintf("Serving from cache %v", h.Pattern()))
		_, _ = fmt.Fprint(w, cachedResponse)
		return
	}

	statistics := application.GetGeneralStatistics(h.repository, params)
	responseBody := []interface{}{params, statistics}
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(responseBody)

	h.log.Info("Statistics request", r.RequestURI)
	if _, err := fmt.Fprint(w, string(jsonData)); err != nil {
		h.log.Error("Failed to write response", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	err = h.cache.Set(r.Context(), h.Pattern()+fmt.Sprintf("%#v", params), jsonData, 60*time.Second)
	if err != nil {
		h.log.Error("Failed to cache response", err)
		return
	}
}
