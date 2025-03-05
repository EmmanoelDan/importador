package model

type Sku struct {
	SkuId          string `gorm:"primaryKey;column:skuid"`
	SkuName        string `gorm:"column:skuname"`
	AvailabilityId string `gorm:"column:availabilityid"`
}
