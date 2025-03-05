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

func (repo *BillingRepository) FindAllWithRelations(page, pageSize int) ([]model.Billing, error) {
	var billings []model.Billing
	err := repo.DB.
		Debug().
		Preload("Customer").
		Preload("Customer.Partner").
		Preload("Product").
		Preload("Product.Sku").
		Preload("Entitlement").
		Limit(pageSize).               // Limita o número de resultados
		Offset((page - 1) * pageSize). // Define o offset baseado na página
		Find(&billings).Error

	return billings, err
}
