package models

import (
	"time"

	"gorm.io/gorm"
)

type Stock struct {
	gorm.Model

	Ticker      string    `gorm:"primaryKey;column:ticker" json:"ticker"`
	Price       string    `gorm:"column:price;not null" json:"price"`
	Company     string    `gorm:"column:company;not null" json:"company"`
	Brokerage   string    `gorm:"column:brokerage;not null" json:"brokerage"`
	Action      string    `gorm:"column:action;not null" json:"action"`
	RatingFrom  string    `gorm:"column:rating_from;not null" json:"rating_from"`
	RatingTo    string    `gorm:"column:rating_to;not null" json:"rating_to"`
	TargetFrom  string    `gorm:"column:target_from;not null" json:"target_from"`
	TargetTo    string    `gorm:"column:target_to;not null" json:"target_to"`
	Time        time.Time `gorm:"column:time;not null" json:"time"`
	Recommended bool      `json:"recommended" gorm:"-"`
}
