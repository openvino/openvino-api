package model

import "time"

// TokenWinerie - Base GORM Model
type TokenWinerie struct {
	ID              string    `json:"id" gorm:"primaryKey"`
	Image           string    `json:"image"`
	BottleImage     string    `json:"bottle_image"`
	TokenIcon       string    `json:"token_icon"`
	CrowSaleAddress string    `json:"crow_sale_address"`
	TokenAddress    string    `json:"token_address"`
	RedeemDate      time.Time `json:"redeem_date"`
	Metrics         Metrics   `json:"metrics" gorm:"embedded"`
	Year            int       `json:"year"`
	Open            bool      `json:"open"`
	Stage           string    `json:"stage"`
	WinerieID       string
	Winerie         Winerie
}

type Metrics struct {
	GrapeCultivation int `json:"grape_cultivation"`
	WineProduction   int `json:"wine_production"`
	Packaging        int `json:"packaging"`
	Logistics        int `json:"logistics"`
	Administration   int `json:"administration"`
	Marketing        int `json:"marketing"`
}
