package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type PartnerRepository struct {
	DB *gorm.DB
}

func (repo *PartnerRepository) Create(partner *model.Partner) error {
	return repo.DB.Create(partner).Error
}