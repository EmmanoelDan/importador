package model

type Subscription struct {
	SubscriptionId        string `gorm:"primaryKey;column:subscriptionid"`
	CustomerId            string `gorm:"column:customerid"`
	ProductId             string `gorm:"column:productid"`
	SubscriptionStartDate string `gorm:"column:subscriptionstartdate"`
}