package application

import (
	"fmt"
	"time"
)

type MostActiveCategoryByTenders struct {
	Category string `json:"category"`
	Count    int    `json:"count"`
}
type MostActiveCategoryByPriceSum struct {
	Category string  `json:"category"`
	Sum      float64 `json:"sum"`
}
type CategorySumsCount struct {
	Category string  `json:"category"`
	Sum      float64 `json:"sum"`
	Count    int     `json:"count"`
}
type TendersPerMonth struct {
	TendersCount int `json:"tenders_count"`
	Year         int `json:"year"`
	Month        int `json:"month"`
}
type DiagramDataPerMonth struct {
	TendersCount int     `json:"tenders_count"`
	TendersSum   float64 `json:"tenders_sum,omitempty"`
	Year         int     `json:"year"`
	Month        int     `json:"month"`
}

type GeneralStatistics struct {
	ActiveTenders                    int                          `json:"active_tenders"`
	MostActiveCategoryByTenders      MostActiveCategoryByTenders  `json:"most_active_category_by_tenders"`
	MostActiveCategoryByPriceSum     MostActiveCategoryByPriceSum `json:"most_active_category_by_price_sum"`
	CategorySumsCounts               []CategorySumsCount          `json:"category_sums_counts"`
	MonthsWithMoreTendersThanAverage []TendersPerMonth            `json:"months_with_more_tenders_than_average"`
	DiagramData                      []DiagramDataPerMonth        `json:"diagram_data"`
	TimeSpent                        time.Duration                `json:"time_spent"`
}

type Params struct {
	ShowSum      bool
	SumRangeFrom int
	SumRangeTo   int
	CategoryCode string
	KeyWords     []string
}

func GetGeneralStatistics(repo StatisticsRepository, params Params) GeneralStatistics {
	fmt.Println("GetGeneralStatistics start", time.Now().Format("15:04:05.000"))
	var statistics GeneralStatistics
	start := time.Now()
	//var wg sync.WaitGroup
	//wg.Add(6)
	//go func() {
	//	statistics.ActiveTenders = repo.ActiveTenders()
	//	fmt.Println("ActiveTenders", time.Since(start))
	//	wg.Done()
	//}()
	//
	//go func() {
	//	statistics.MostActiveCategoryByTenders.CategoryCode,
	//		statistics.MostActiveCategoryByTenders.Count = repo.MostActiveCategoryByTenders()
	//	fmt.Println("MostActiveCategoryByTenders", time.Since(start))
	//	wg.Done()
	//}()
	//
	//go func() {
	//	statistics.MostActiveCategoryByPriceSum.CategoryCode,
	//		statistics.MostActiveCategoryByPriceSum.Sum = repo.MostActiveCategoryByPriceSum()
	//	fmt.Println("MostActiveCategoryByPriceSum", time.Since(start))
	//	wg.Done()
	//}()
	//
	//go func() {
	//	statistics.CategorySumsCounts = repo.CategorySumsCounts()
	//	fmt.Println("CategorySumsCounts", time.Since(start))
	//	wg.Done()
	//}()
	//
	//go func() {
	//	statistics.MonthsWithMoreTendersThanAverage = repo.MonthsWithMoreTendersThanAverage()
	//	fmt.Println("MonthsWithMoreTendersThanAverage", time.Since(start))
	//	wg.Done()
	//}()
	//
	//go func() {
	//	statistics.DiagramData = repo.DiagramByDate()
	//	fmt.Println("DiagramByDate", time.Since(start), statistics.DiagramData)
	//	wg.Done()
	//}()
	//wg.Wait()
	func() {
		statistics.DiagramData = repo.DiagramByDate(params)
		fmt.Println("DiagramByDate", time.Since(start), statistics.DiagramData)
	}()
	statistics.TimeSpent = time.Duration(time.Since(start).Seconds())
	fmt.Println(start, statistics.TimeSpent)
	//fmt.Println(statistics)
	return statistics
}

func GetPersonalStatistics(userId string) string {
	statistics := "Personal statistics for user " + userId
	return statistics
}

func TestResponseTime() string {
	time.Sleep(400 * time.Millisecond)
	return "success " + fmt.Sprint(time.Now().Format("15:04:05.000"))
}
