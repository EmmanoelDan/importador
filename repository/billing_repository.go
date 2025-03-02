package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type BillingRepository struct {
	DB *gorm.DB
}

func (repo *BillingRepository) Create(billing *model.Billing) error {
	return repo.DB.Create(billing).Error
}
