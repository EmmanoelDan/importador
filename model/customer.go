package model

type Customer struct {
	CustomerId         string `gorm:"primaryKey; column:customerid"`
	CustomerName       string `gorm:"column:customername"`
	CustomerDomainName string `gorm:"column:customerdomainname"`
	CustomerCountry    string `gorm:"column:customercountry"`
	MpnId              string `gorm:"column:mpnid"`
	Tier2MpnId         string `gorm:"column:tier2mpnid"`
}