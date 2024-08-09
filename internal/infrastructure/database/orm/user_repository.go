package orm

import (
	"git.b4i.kz/b4ikz/tenderok-analytics/internal/domain"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (u *UserRepositoryImpl) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	var userModel User
	result := u.DB.First(&userModel, id)
	if result.Error != nil {
		return &domain.User{}, result.Error
	}
	err := copier.Copy(&user, &userModel)
	return &user, err
}
