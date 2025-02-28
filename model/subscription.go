package model

type Subscription struct {
	SubscriptionId          string `gorm:"primaryKey;column:subscriptionid"`
	CustomerId              string `gorm:"column:customerid"`
	ProductId               string `gorm:"column:productid"`
	SubscriptionDescription string `gorm:"column:subscriptiondescription"`
}