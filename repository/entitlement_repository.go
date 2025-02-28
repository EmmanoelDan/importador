package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type EntitlementRepository struct {
	DB *gorm.DB
}

func (repo *EntitlementRepository) Create(entitlement *model.Entitlement) error {
    return repo.DB.Create(entitlement).Error
}