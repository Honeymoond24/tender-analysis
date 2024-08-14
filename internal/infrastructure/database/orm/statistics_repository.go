package orm

import "gorm.io/gorm"

type StatisticsRepository struct {
	DB *gorm.DB
}

func (s *StatisticsRepository) MostActiveCategoryByTenders() (string, int) {
	//TODO implement me
	panic("implement me")
}

func (s *StatisticsRepository) MostActiveCategoryByPriceSum() (string, float64) {
	//TODO implement me
	panic("implement me")
}

func (s *StatisticsRepository) ActiveTenders() int64 {
	var count int64
	s.DB.Model(&Tender{}).Count(&count)
	return count
}
