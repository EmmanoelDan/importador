package model

import (
	"time"

	"gorm.io/datatypes"
)

type Billing struct {
	BillingId              string         `gorm:"primaryKey;column:billingid;autoIncrement"`
	CustomerId             string         `gorm:"column:customerid"`
	ProductId              string         `gorm:"column:productid"`
	EntitlementId          string         `gorm:"column:entitlementid"`
	InvoiceNumber          string         `gorm:"column:invoicenumber"`
	ChargeStartDate        time.Time      `gorm:"column:chargestartdate;type:date"`
	ChargeEndDate          time.Time      `gorm:"column:chargeenddate;type:date"`
	UsageDate              time.Time      `gorm:"column:usagedate;type:date"`
	MeterId                string         `gorm:"column:meterid"`
	MeterType              string         `gorm:"column:metertype"`
	MeterCategory          string         `gorm:"column:metercategory"`
	MeterSubCategory       string         `gorm:"column:metersubcategory"`
	MeterName              string         `gorm:"column:metername"`
	MeterRegion            string         `gorm:"column:meterregion"`
	Unit                   string         `gorm:"column:unit"`
	ResourceLocation       string         `gorm:"column:resourcelocation"`
	ConsumedService        string         `gorm:"column:consumedservice"`
	ResourceGroup          string         `gorm:"column:resourcegroup"`
	ResourceURI            string         `gorm:"column:resourceuri"`
	ChargeType             string         `gorm:"column:chargetype"`
	UnitPrice              float64        `gorm:"column:unitprice"`
	Quantity               float64        `gorm:"column:quantity"`
	UnitType               string         `gorm:"column:unittype"`
	BillingPreTaxTotal     float64        `gorm:"column:billingpretaxtotal"`
	BillingCurrency        string         `gorm:"column:billingcurrency"`
	PricingPreTaxTotal     float64        `gorm:"column:pricingpretaxtotal"`
	PricingCurrency        string         `gorm:"column:pricingcurrency"`
	ServiceInfo1           string         `gorm:"column:serviceinfo1"`
	ServiceInfo2           string         `gorm:"column:serviceinfo2"`
	Tags                   datatypes.JSON `gorm:"column:tags;type:json"`
	AdditionalInfo         datatypes.JSON `gorm:"column:additionalinfo;type:json"`
	EffectiveUnitPrice     float64        `gorm:"column:effectiveunitprice"`
	PCToBCExchangeRate     float64        `gorm:"column:pctobcexchangerate"`
	PCToBCExchangeRateDate time.Time      `gorm:"column:pctobcexchangeratedate;type:date"`

	Customer    Customer    `gorm:"foreignKey:CustomerId;references:CustomerId"`
	Product     Product     `gorm:"foreignKey:ProductId;references:ProductId"`
	Entitlement Entitlement `gorm:"foreignKey:EntitlementId;references:EntitlementId"`
}
