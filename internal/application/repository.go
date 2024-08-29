package application

type StatisticsRepository interface {
	ActiveTenders() int
	MostActiveCategoryByTenders() (string, int)
	MostActiveCategoryByPriceSum() (string, float64)
	CategorySumsCounts() []CategorySumsCount
	MonthsWithMoreTendersThanAverage() []TendersPerMonth
	DiagramByDate(params Params) []DiagramDataPerMonth
}

type PersonalStatisticsRepository interface {
}
