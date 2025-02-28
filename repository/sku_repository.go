package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type SkuRepository struct {
	DB *gorm.DB
}

func (repo *SkuRepository) Create(sku *model.Sku) error {
    return repo.DB.Create(sku).Error
}