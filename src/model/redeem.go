package model

import "github.com/jinzhu/gorm"

// Sale - Base GORM Model
type RedeemInfo struct {
	gorm.Model
	CustomerId     string `json:"customer_id"`
	Customer       User   `json:"customer" gorm:"foreignKey:CustomerId"`
	Year           string `json:"year"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	CountryId      uint   `json:"country_id"`
	ProvinceId     uint   `json:"province_id"`
	Zip            string `json:"zip"`
	TelegramId     string `json:"telegram_id"`
	Amount         uint   `json:"amount"`
	Signature      string `json:"signature"`
	BurnTxHash     string `json:"burn_tx_hash"`
	ShippingTxHash string `json:"shipping_tx_hash"`
}

// Sale - Base GORM Model
type ShippingCost struct {
	gorm.Model
	CountryId  uint    `json:"country_id"`
	ProvinceId uint    `json:"province_id"`
	Amount     uint    `json:"amount"`
	Cost       float32 `json:"cost"`
}
