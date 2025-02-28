package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type Billing struct {
	DB *gorm.DB
}

func (repo *Billing) Create(billing *model.Billing) error {
    return repo.DB.Create(billing).Error
}