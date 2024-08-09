package application

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/domain"
)

type UserRepository interface {
	GetUserByID(id uint) (*domain.User, error)
}

type TenderRepository interface {
	GetOne(id int) (*domain.Tender, error)
}

type CompanyRepository interface {
	GetOne(id int) (*domain.Company, error)
}

type GeneralStatistics interface {
	ActiveTenders() int64
	MostActiveCategoryByTenders() (string, int)
	MostActiveCategoryByPriceSum() (string, float64)
}
