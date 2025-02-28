package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type MeterRepository struct {
	DB *gorm.DB
}

func (repo *MeterRepository) Create(meter *model.Meter) error {
    return repo.DB.Create(meter).Error
}