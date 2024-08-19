package use_cases

import (
	"fmt"
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/application"
)

type MostActiveCategoryByTenders struct {
	Category string
	Count    int
}
type MostActiveCategoryByPriceSum struct {
	Category string
	Sum      float64
}

type GeneralStatistics struct {
	ActiveTenders                    int64
	MostActiveCategoryByTenders      MostActiveCategoryByTenders
	MostActiveCategoryByPriceSum     MostActiveCategoryByPriceSum
	CategorySumsCounts               [][]interface{}
	MonthsWithMoreTendersThanAverage [][]int
}

func GetGeneralStatistics(repo application.StatisticsRepository) GeneralStatistics {
	var statistics GeneralStatistics
	statistics.ActiveTenders = repo.ActiveTenders()
	statistics.MostActiveCategoryByTenders.Category,
		statistics.MostActiveCategoryByTenders.Count = repo.MostActiveCategoryByTenders()
	statistics.MostActiveCategoryByPriceSum.Category,
		statistics.MostActiveCategoryByPriceSum.Sum = repo.MostActiveCategoryByPriceSum()
	statistics.CategorySumsCounts = repo.CategorySumsCounts()
	statistics.MonthsWithMoreTendersThanAverage = repo.MonthsWithMoreTendersThanAverage()
	fmt.Println(statistics)
	return statistics
}

func GetPersonalStatistics(userId string) string {
	statistics := "Personal statistics for user " + userId
	return statistics
}
