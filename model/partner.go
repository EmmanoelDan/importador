package model

type Partner struct {
	PartnerId   string `gorm:"primaryKey;column:partnerid"`
	PartnerName string `gorm:"column:partnername"`
	MpnId       string `gorm:"column:mpnid"`
	Tier2MpnId  string `gorm:"column:tier2mpnid"`
}
