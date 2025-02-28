package repository

import (
	"github.com/EmmanoelDan/importador/model"
	"gorm.io/gorm"
)

type SubscriptionRepository struct {
	DB *gorm.DB
}

func (repo *SubscriptionRepository) Create(subscription *model.Subscription) error {
    return repo.DB.Create(subscription).Error
}