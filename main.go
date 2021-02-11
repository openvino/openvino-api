package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/openvino/openvino-api/src/config"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
)

func main() {

	log.Println("----------------------------")
	log.Println("1. Getting config variables...")
	config.Config = config.New()

	log.Println("2. Setting up router...")
	router := NewRouter()

	var err error

	log.Println("3. Connecting to database...")
	repository.DB, err = repository.SetupDB(config.Config.Database)
	if err != nil {
		log.Panicf("Unable to connect to database: %s", err.Error())
	}

	repository.Eth, err = repository.SetupETH(config.Config.Ethereum)
	if err != nil {
		log.Panicf("Unable to connect to infura: %s", err.Error())
	}

	log.Println("4. Migrating database model...")
	repository.DB.AutoMigrate(
		&model.Sale{}, &model.SensorRecord{}, &model.User{},
		&model.Task{}, &model.Tools{}, &model.Chemicals{},
		&model.RedeemInfo{}, &model.ShippingCost{})
	defer repository.DB.Close()

	log.Println("----------------------------")

	log.Printf("Serving application at PORT: %s", config.Config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Config.Port), router))
}
