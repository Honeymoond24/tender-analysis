package use_cases

import "git.b4i.kz/b4ikz/tenderok-analytics/internal/application"

func GetGeneralStatistics(repo application.StatisticsRepository) string {

	statistics := "General statistics"
	return statistics
}

func GetPersonalStatistics(userId string) string {
	statistics := "Personal statistics for user " + userId
	return statistics
}
