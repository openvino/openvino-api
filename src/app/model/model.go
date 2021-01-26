package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SensorData struct {
	Hash string `gorm:"primary_key" json:"hash"`

	Timestamp *time.Time `json:"timestamp"`
	SensorID  string     `json:"sensor_id"`

	Humidity2   float64 `json:"humidity2"`
	Humidity1   float64 `json:"humidity1"`
	Humidity05  float64 `json:"humidity05"`
	Humidity005 float64 `json:"humidity005"`

	WindVelocity  float64 `json:"wind_velocity"`
	WindGust      float64 `json:"wind_gust"`
	WindDirection float64 `json:"wind_direction"`
	Pressure      float64 `json:"pressure"`
	Rain          float64 `json:"rain"`
	Temperature   float64 `json:"temperature"`
	Humidity      float64 `json:"humidity"`
	IrradianceIR    float64 `json:"irradiance_ir"`
	IrradianceUV    float64 `json:"irradiance_uv"`
	IrradianceVI    float64 `json:"irradiance_vi"`
}

type Root struct {
	Root   string `gorm:"primary_key" json:"root"`
	TxHash string `json:"tx_hash"`
}

type Buyer struct {
	PublicKey string `gorm:"primary_key" json:"public_key"`
	Email     string `json:"email"`
	Amount    int    `json:"amount"`
}

type Task struct {
	Hash    string     `gorm:"primary_key" json:"hash"`
	Employe string     `json:"employe"`
	Task    string     `json:"task"`
	Ini     *time.Time `json:"ini"`
	End     *time.Time `json:"end"`
	Notes   string     `json:"notes"`
	Status  int        `gorm:"DEFAULT:0" json:"status"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&SensorData{})
	db.AutoMigrate(&Root{})
	db.AutoMigrate(&Buyer{})
	db.AutoMigrate(&Task{})
	return db
}
