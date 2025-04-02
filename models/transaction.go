package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID   uint    `json:"user_id" gorm:"not null"` //to locate who the user is 
	Type     string  `json:"type" gorm:"not null"` //type of transaction -  of two categories, either income or an expense
	Category string  `json:"category" gorm:"not null"`
	Amount   float64 `json:"amount" gorm:"not null"`
	Note 	string  `json:"note"`
}
