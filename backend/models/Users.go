package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	FirtsName string `gorm:"column:first_name;not null"`
	LastName  string `gorm:"column:last_name;not null"`
	Email     string `gorm:"column:email;not null"`
	Password  string `gorm:"column:password;not null"`
}
