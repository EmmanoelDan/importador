package model

type User struct {
	UserId   string `gorm:"primaryKey;column:userid;autoIncrement"`
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}
