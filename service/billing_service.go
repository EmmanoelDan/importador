package service

import (
	"github.com/EmmanoelDan/importador/model"
	"github.com/EmmanoelDan/importador/repository"
)

type BillingService struct {
	BillingRepo *repository.BillingRepository
}

func NewBillingService(billingRepo *repository.BillingRepository) *BillingService {
	return &BillingService{BillingRepo: billingRepo}
}

func (s *BillingService) FindAllWithRelations(page, pageSize int) ([]model.Billing, error) {
	return s.BillingRepo.FindAllWithRelations(page, pageSize)
}
