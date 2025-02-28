package model

type Meter struct {
	MeterId          string `gorm:"primaryKey;column:meterid"`
	MeterType        string `gorm:"column:metertype"`
	MeterCategory    string `gorm:"column:metercategory"`
	MeterSubCategory string `gorm:"column:metersubcategory"`
	MeterName        string `gorm:"column:metername"`
	MeterRegion      string `gorm:"column:meterregion"`
	Unit             string `gorm:"column:unit"`
}