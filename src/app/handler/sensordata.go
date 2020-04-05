package handler

import (
	"fmt"
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/openvino/openvino-api/src/app/model"
)

func GetSensorData(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
}

func GetSensorDataDay(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	day := r.URL.Query().Get("day")
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	sensordata := []model.SensorData{}
	fmt.Println("Hello, Im a day")

	db.Where("DAY(timestamp) = ? AND MONTH(timestamp) = ? AND YEAR(timestamp) = ?", day, month, year).Find(&sensordata)
	respondJSON(w, http.StatusOK, sensordata)
	
}

func GetSensorDataMonth(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	sensordata := []model.SensorData{}
	fmt.Println("Hello, Im a month")
	db.Where("MONTH(timestamp) = ? AND YEAR(timestamp) = ?", month, year).Find(&sensordata)
	respondJSON(w, http.StatusOK, sensordata)

}

func GetSensorDataYear(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	year := r.URL.Query().Get("year")

	sensordata := []model.SensorData{}
	fmt.Println("Hello, Im a nigger")

	db.Where("YEAR(timestamp) = ?", year).Find(&sensordata)
	respondJSON(w, http.StatusOK, sensordata)

}

func CreateSensorData(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	sensordata := model.SensorData{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&sensordata); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&sensordata).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	
	respondJSON(w, http.StatusCreated, sensordata)

}
