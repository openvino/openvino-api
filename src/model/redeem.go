package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Sale - Base GORM Model
type RedeemInfo struct {
	ID             string `gorm:"primary_key"`
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
	CreatedAt      time.Time
	UpdatedAt      time.Time
	WinerieID      string
	Winerie        Winerie
	Status string `json:"redeem_status"`
	Watched bool `json:"watched"`
	City string `json:"city"`
	Phone string `json:"phone"`


}

// Sale - Base GORM Model
type ShippingCost struct {
	gorm.Model
	CountryId   uint    `json:"country_id"`
	ProvinceId  uint    `json:"province_id"`
	BaseCost    float64 `json:"base_cost"`
	CostPerUnit float64 `json:"cost_per_unit"`
}
