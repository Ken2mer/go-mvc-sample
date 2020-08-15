package model

type User struct {
	UserID uint `gorm:"primary_key"`
	Name   string
}
