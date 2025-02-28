package model

type Entitlement struct {
	EntitlementId                 string `gorm:"primaryKey;column:entitlementid"`
	EntitlementDescription        string `gorm:"column:entitlementdescription"`
	PartnerEarnedCreditPercentage string `gorm:"column:partnerearnedcreditpercentage"`
	CreditPercentage              string `gorm:"column:creditpercentage"`
	CreditType                    string `gorm:"column:credittype"`
	BenefitOrderId                string `gorm:"column:benefitorderid"`
	BenefitId                     string `gorm:"column:benefitid"`
	BenefitType                   string `gorm:"column:benefittype"`
}