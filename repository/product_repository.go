package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func (repo *ProductRepository) Create(product *model.Product) error {
	return repo.DB.Create(product).Error
}