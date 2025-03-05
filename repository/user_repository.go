package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) Create(user *model.User) error {
	return repo.DB.Create(user).Error
}

func (repo *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User

	err := repo.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
