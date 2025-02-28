package model

type SKU struct {
	SkuId          string `gorm:"primaryKey;column:skuid"`
	ProductId      string `gorm:"column:productid"`
	SkuName        string `gorm:"column:skuname"`
	AvailabilityId string `gorm:"column:availabilityid"`
}