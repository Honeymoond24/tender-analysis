package controllers

import (
	"fmt"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application/use_cases"
	"net/http"
)

func GeneralStatistics(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI, "user id - ", r.PathValue("id"))
	statistics := use_cases.GetGeneralStatistics()
	_, err := w.Write([]byte(statistics))
	if err != nil {
		return
	}
}
