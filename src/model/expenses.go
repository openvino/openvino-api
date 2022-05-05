package model

import (
	"time"
)

// Expenses - Base GORM Model
type Expense struct {
	Hash        string     `json:"hash" gorm:"primaryKey"`
	Token       uint       `json:"token_id"`
	Timestamp   *time.Time `json:"timestamp"`
	TypeId      uint       `json:"expense_id"`
	Description string     `json:"description"`
	Value       float32    `json:"value"`
	WinerieID   int
	Winerie     Winerie
}

type Token struct {
	Id     uint    `gorm:"primaryKey" json:"id"`
	Amount float32 `json:"amount"`
}
