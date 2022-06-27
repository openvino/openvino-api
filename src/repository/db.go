package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/openvino/openvino-api/src/config"

	// Import mysql driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB - Global DB variable
var DB *gorm.DB

// SetupDB - Sets up a Database Connection
func SetupDB(config config.DatabaseConfig) (*gorm.DB, error) {

	var dbHost string = config.Host
	var dbPort string = config.Port
	var dbName string = config.DatabaseName
	var dbUser string = config.Username
	var dbPassword string = config.Password

	db, dbError := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName))
	if dbError != nil {
		return nil, dbError
	}

	db.DB().SetMaxIdleConns(0)

	return db, nil
}
