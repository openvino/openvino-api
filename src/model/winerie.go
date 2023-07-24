package model

import "time"

type Winerie struct {
	ID           string     `gorm:"primary_key"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `sql:"index"`
	Name         string     `json:"name"`
	Website      string     `json:"website"`
	Image        string     `json:"image"`
	PrimaryColor string     `json:"primary_color"`
	Secret       string     `json:"secret"`
	Email       string     `json:"email"`
	PublicKey string `json:"public_key"`
	IsAdmin string `json:"isAdmin"`
}
