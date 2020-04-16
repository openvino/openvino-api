package handler

import (
	"time"
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/openvino/openvino-api/src/app/model"
)

func GetSensorDataWrong(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusBadRequest, "Malformed query")
}

func GetSensorDataDay(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	day := r.URL.Query().Get("day")
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	sensordata := []model.SensorData{}

	db.Where("DAY(timestamp) = ? AND MONTH(timestamp) = ? AND YEAR(timestamp) = ?", day, month, year).Find(&sensordata)
	respondJSON(w, http.StatusOK, sensordata)
	
}

func GetSensorDataMonth(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	sensordata := []model.SensorData{}

	db.Where("MONTH(timestamp) = ? AND YEAR(timestamp) = ?", month, year).Find(&sensordata)
	respondJSON(w, http.StatusOK, sensordata)

}

func GetSensorDataYear(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	year := r.URL.Query().Get("year")

	sensordata := []model.SensorData{}

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

	now := time.Now()
	sensordata.Timestamp = &now

	if err := db.FirstOrCreate(&sensordata).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusCreated, sensordata)

}