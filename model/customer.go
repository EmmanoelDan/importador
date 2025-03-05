package model

type Customer struct {
	CustomerId         string `gorm:"primaryKey; column:customerid"`
	CustomerName       string `gorm:"column:customername"`
	CustomerDomainName string `gorm:"column:customerdomainname"`
	CustomerCountry    string `gorm:"column:customercountry"`
	PartnerId          string `gorm:"column:partnerid"`

	Partner Partner `gorm:"foreignKey:PartnerId;references:PartnerId"`
}
