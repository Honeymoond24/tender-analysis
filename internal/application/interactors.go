package application

import (
	"fmt"
	"sync"
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
	Count    int64   `json:"count"`
}
type TendersPerMonth struct {
	TendersCount int64 `json:"tenders_count"`
	Year         int   `json:"year"`
	Month        int   `json:"month"`
}

type GeneralStatistics struct {
	ActiveTenders                    int64                        `json:"active_tenders"`
	MostActiveCategoryByTenders      MostActiveCategoryByTenders  `json:"most_active_category_by_tenders"`
	MostActiveCategoryByPriceSum     MostActiveCategoryByPriceSum `json:"most_active_category_by_price_sum"`
	CategorySumsCounts               []CategorySumsCount          `json:"category_sums_counts"`
	MonthsWithMoreTendersThanAverage []TendersPerMonth            `json:"months_with_more_tenders_than_average"`
	TimeSpent                        time.Duration                `json:"time_spent"`
}

func GetGeneralStatistics(repo StatisticsRepository) GeneralStatistics {
	fmt.Println("GetGeneralStatistics start", time.Now().Format("15:04:05.000"))
	var statistics GeneralStatistics
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(5)
	go func() {
		statistics.ActiveTenders = repo.ActiveTenders()
		fmt.Println("ActiveTenders", time.Since(start))
		wg.Done()
	}()

	go func() {
		statistics.MostActiveCategoryByTenders.Category,
			statistics.MostActiveCategoryByTenders.Count = repo.MostActiveCategoryByTenders()
		fmt.Println("MostActiveCategoryByTenders", time.Since(start))
		wg.Done()
	}()

	go func() {
		statistics.MostActiveCategoryByPriceSum.Category,
			statistics.MostActiveCategoryByPriceSum.Sum = repo.MostActiveCategoryByPriceSum()
		fmt.Println("MostActiveCategoryByPriceSum", time.Since(start))
		wg.Done()
	}()

	go func() {
		statistics.CategorySumsCounts = repo.CategorySumsCounts()
		fmt.Println("CategorySumsCounts", time.Since(start))
		wg.Done()
	}()

	go func() {
		statistics.MonthsWithMoreTendersThanAverage = repo.MonthsWithMoreTendersThanAverage()
		fmt.Println("MonthsWithMoreTendersThanAverage", time.Since(start))
		wg.Done()
	}()
	wg.Wait()
	statistics.TimeSpent = time.Duration(time.Since(start).Seconds())
	fmt.Println(start, statistics.TimeSpent)
	//fmt.Println(statistics)
	return statistics
}

func GetPersonalStatistics(userId string) string {
	statistics := "Personal statistics for user " + userId
	return statistics
}

var counter int

func TestResponseTime() string {
	counter++
	localCounter := counter
	start := time.Now()
	fmt.Println("TestResponseTime", localCounter, time.Now())
	time.Sleep(5 * time.Second)
	fmt.Println("TestResponseTime", localCounter, time.Since(start))
	return "success"
}
