package model

type Product struct {
	ProductId               string `gorm:"primaryKey;column:productid"`
	SkuId                   string `gorm:"column:skuid"`
	ProductName             string `gorm:"column:productname"`
	PublisherName           string `gorm:"column:publishername"`
	PublisherId             string `gorm:"column:publisherid"`
	SubscriptionId          string `gorm:"column:subscriptionid"`
	SubscriptionDescription string `gorm:"column:subscriptiondescription"`

	Sku Sku `gorm:"foreignKey:SkuId; references:SkuId;"`
}
