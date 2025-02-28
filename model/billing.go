package model

import "gorm.io/datatypes"

type Billing struct {
	BillingId              string          `gorm:"primaryKey;column:billingid;autoIncrement"`
	PartnerId              string          `gorm:"column:partnerid"`
	SubscriptionId         string          `gorm:"column:subscriptionid"`
	MeterId                string          `gorm:"column:meterid"`
	EntitlementId          string          `gorm:"column:entitlementid"`
	InvoiceNumber          string          `gorm:"column:invoicenumber"`
	ChargeStartDate        string          `gorm:"column:chargestartdate"`
	ChargeEndDate          string          `gorm:"column:chargeenddate"`
	UsageDate              string          `gorm:"column:usagedate"`
	ResourceLocation       string          `gorm:"column:resourcelocation"`
	ConsumedService        string          `gorm:"column:consumedservice"`
	ResourceGroup          string          `gorm:"column:resourcegroup"`
	ResourceURI            string          `gorm:"column:resourceuri"`
	ChargeType             string          `gorm:"column:chargetype"`
	UnitPrice              string          `gorm:"column:unitprice"`
	Quantity               string          `gorm:"column:quantity"`
	UnitType               string          `gorm:"column:unittype"`
	BillingPreTaxTotal     string          `gorm:"column:billingpretaxtotal"`
	BillingCurrency        string          `gorm:"column:billingcurrency"`
	PricingPreTaxTotal     string          `gorm:"column:pricingpretaxtotal"`
	PricingCurrency        string          `gorm:"column:pricingcurrency"`
	ServiceInfo1           string          `gorm:"column:serviceinfo1"`
	ServiceInfo2           string          `gorm:"column:serviceinfo2"`
	Tags                   datatypes.JSON      `gorm:"column:tags;type:json"`
	AdditionalInfo         datatypes.JSON `gorm:"column:additionalinfo;type:json"`
	EffectiveUnitPrice     string          `gorm:"column:effectiveunitprice"`
	PCToBCExchangeRate     string          `gorm:"column:pctobcexchangerate"`
	PCToBCExchangeRateDate string          `gorm:"column:pctobcexchangeratedate"`
}