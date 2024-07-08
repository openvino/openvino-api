package model

import (
	"time"
)

type RedeemLog struct {
	ID             string `gorm:"primary_key"`
	CustomerId     string `json:"customer_id"`
	Customer       User   `json:"customer" gorm:"foreignKey:CustomerId"`
	Year           string `json:"year"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	CountryId      string   `json:"country_id"`
	ProvinceId     string   `json:"province_id"`
	Zip            string `json:"zip"`
	TelegramId     string `json:"telegram_id"`
	Amount         uint   `json:"amount"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	WinerieID      string
	Winerie        Winerie
	City string `json:"city"`
	Phone string `json:"phone"`
	Signature      string `json:"signature"`
	BurnTxHash     string `json:"burn_tx_hash"`
	ShippingTxHash string `json:"shipping_tx_hash"`
	ErrorMessage string `json:"error_message"`
	ShippingPaidStatus string `json:"shipping_paid_status"`
	Pickup string `json:"pickup"`
}

