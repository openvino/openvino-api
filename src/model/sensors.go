package model

import "time"

// SensorRecord - Base GORM Model
type SensorRecord struct {
	Hash      string     `gorm:"primary_key" json:"hash"`
	Timestamp *time.Time `json:"timestamp"`

	SensorId string `json:"sensor_id"`

	Humidity2   float64 `json:"humidity2"`
	Humidity1   float64 `json:"humidity1"`
	Humidity05  float64 `json:"humidity05"`
	Humidity005 float64 `json:"humidity005"`

	WindVelocity  float64 `json:"wind_velocity"`
	WindGust      float64 `json:"wind_gust"`
	WindDirection float64 `json:"wind_direction"`

	Pressure float64 `json:"pressure"`

	Rain float64 `json:"rain"`

	Temperature float64 `json:"temperature"`

	Humidity float64 `json:"humidity"`

	IrradianceIR float64 `json:"irradiance_ir"`
	IrradianceUV float64 `json:"irradiance_uv"`
	IrradianceVI float64 `json:"irradiance_vi"`
	WinerieID    string
	Winerie      Winerie
}
