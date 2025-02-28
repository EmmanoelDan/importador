package model

type Partner struct {
	PartnerId   string `gorm:"primaryKey;column:partnerid"`
	PartnerName string `gorm:"column:partnername"`
}