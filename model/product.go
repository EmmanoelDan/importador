package model

type Product struct {
	ProductId     string `gorm:"primaryKey;column:productid"`
	ProductName   string `gorm:"column:productname"`
	PublisherName string `gorm:"column:publishername"`
	PublisherId   string `gorm:"column:publisherid"`
}