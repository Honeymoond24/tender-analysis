package application

type StatisticsRepository interface {
	ActiveTenders() int64
	MostActiveCategoryByTenders() (string, int)
	MostActiveCategoryByPriceSum() (string, float64)
	CategorySumsCounts() [][]interface{}
	MonthsWithMoreTendersThanAverage() [][]int
}
