package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func (repo *CustomerRepository) Create(customer *model.Customer) error {
	return repo.DB.Create(customer).Error
}