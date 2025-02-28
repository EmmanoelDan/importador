package model

type Entitlement struct {
	EntitlementId                 string `gorm:"primaryKey"`
	EntitlementDescription        string `gorm:"column:entitlementdescription"`
	PartnerEarnedCreditPercentage int16  `gorm:"column:partnerEarnedCreditPercentage"`
	CreditPercentage              int16  `gorm:"column:creditPercentage"`
	CreditType                    string `gorm:"column:creditType"`
	BenefitOrderId                string `gorm:"column:benefitorder"`
	BenefitId                     string `gorm:"column:benefitid"`
	BenefitType                   string `gorm:"column:benefittype"`
}