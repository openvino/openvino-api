package model

import (
	"time"
)

// Tasks - Base GORM Model
type Task struct {
	Hash          string      `json:"hash" gorm:"primaryKey"`
	PublicKey     string      `json:"public_key"`
	IniTimestamp  *time.Time  `json:"ini_timestamp"`
	IniClaro      string      `json:"ini_claro"`
	IniRow        uint        `json:"ini_row"`
	IniPlant      uint        `json:"ini_plant"`
	EndTimestamp  *time.Time  `json:"end_timestamp"`
	EndClaro      string      `json:"end_claro"`
	EndRow        uint        `json:"end_row"`
	EndPlant      uint        `json:"end_plant"`
	CategoryId    uint        `json:"category_id"`
	TypeId        uint        `json:"task_id"`
	ToolsUsed     []Tools     `json:"tools_used" gorm:"ForeignKey:TaskHash"`
	ChemicalsUsed []Chemicals `json:"chemicals" gorm:"ForeignKey:TaskHash"`
	Notes         string      `json:"notes"`
	WinerieID     int
	Winerie       Winerie
}

// Tools - Base GORM Model
type Tools struct {
	Id       uint   `gorm:"primaryKey json:"tool_id"`
	TaskHash string `gorm:"primaryKey json:"task"`
}

// Chemicals - Base GORM Model
type Chemicals struct {
	Id       uint    `gorm:"primaryKey json:"chemical_id"`
	Amount   float32 `json:"amount"`
	TaskHash string  `gorm:"primaryKey json:"task"`
}
