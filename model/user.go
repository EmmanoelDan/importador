package model

type User struct {
	userId   string `gorm:"primaryKey;column:userid;autoIncrement"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}
