package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type SensorData struct {
	gorm.Model

	Timestamp 		*time.Time   	`gorm:"primary_key" json:"timestamp"`
	SensorID  		string     		`gorm:"primary_key" json:"sensor_id"`

	Humidity2  		int 			`json:"humidity2"`
	Humidity1		int				`json:"humidity1"`
	Humidity05		int				`json:"humidity05"`
	Humidity005		int             `json:"humidity005"`

	WindVelocity 	int				`json:"wind_velocity"`
	WindGust 		int				`json:"wind_gust"`
	WindDirection	int				`json:"wind_direction"`
	Pressure		int				`json:"pressure"`
	Rain 			int				`json:"rain"`
	Temperature		int				`json:"temperature"`
	Humidity		int				`json:"humidity"`
	
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&SensorData{})	
	return db
}
