package orm

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name string
}

type Company struct {
	gorm.Model
	Name string
}

type Tender struct {
	ID            uuid.UUID `gorm:"column:id;type:uuid;primary_key"`
	OpenEndDate   time.Time `gorm:"column:open_end_date"`
	OffersEndDate time.Time `gorm:"column:offers_end_date"`
}
