package models

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	AccountNumber string  `json:"account_number"`
	AccountHolder string  `json:"account_holder"`
	Balance       float64 `json:"balance"`
}
