package model

import "github.com/jinzhu/gorm"

type Winerie struct {
	gorm.Model
	Name         string `json:"name"`
	Website      string `json:"website"`
	Image        string `json:"image"`
	PrimaryColor string `json:"primary_color"`
}
