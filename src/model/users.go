package model

// User - Base GORM Model
type User struct {
	PublicKey string `json:"public_key" gorm:"primary_key"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
