package model

import (
"github.com/jinzhu/gorm"
)

// Sale - Base GORM Model
type Sale struct {
	gorm.Model
	CustomerId string `json:"customer_id"`
	Customer  User	`json:"customer" gorm:"foreignKey:CustomerId"`
	Amount	  int 	`json:"amount"`
}