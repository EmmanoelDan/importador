package model

import (
	"gorm.io/datatypes"
)

type Billing struct {
	BillingId              string         `gorm:"primaryKey;column:billingid"`
	PartnerId              string         `gorm:"column:partnerid"`
	SubscriptionId         string         `gorm:"column:subscriptionid"`
	MeterId                string         `gorm:"column:meterid"`
	EntitlementId          string         `gorm:"column:etitlementid"`
	InvoiceNumber          string         `gorm:"column:invoicenumber"`
	ChargeStartDate        string         `gorm:"column:chargestartdate"`
	ChargeEndDate          string         `gorm:"column:chargesenddate"`
	UsageDate              string         `gorm:"column:usagedate"`
	ResourceLocation       string         `gorm:"column:resourcelocation"`
	ConsumedService        string         `gorm:"column:consumedservice"`
	ResouceGroup           string         `gorm:"column:resoucegroup"`
	ResouceURI             string         `gorm:"column:resouceuri"`
	ChargeType             string         `gorm:"column:chargetype"`
	UnitPrice              float64        `gorm:"column:unitprice"`
	Quantity               float64        `gorm:"column:quantity"`
	UnitType               string         `gorm:"column:unittype"`
	BillingPreTaxTotal     float64        `gorm:"column:billingpretax_total"`
	BillingCurrency        string         `gorm:"column:billingcurrency"`
	PricingPreTaxTotal     string         `gorm:"column:pricingpretax_total"`
	PricingCurrency        string         `gorm:"column:pricingcurrency"`
	ServiceInfo1           string         `gorm:"column:serviceinfo1"`
	ServiceInfo2           string         `gorm:"column:serviceinfo2"`
	Tags                   datatypes.JSON `gorm:"column:tags;type:json"`
	AdditionalInfo         datatypes.JSON `gorm:"column:additionalinfo;type:json"`
	EffectiveUnitPrice     float64        `gorm:"column:effectiveunitprice"`
	PCToBCExchangeRate     int16          `gorm:"column:pctobcexchangerate"`
	PCToBCExchangeRateDate string         `gorm:"column:pctobcexchangeratedate"`
}