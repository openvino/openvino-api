package handler

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/openvino/openvino-api/src/app/model"
)

func GetSensorDataHash(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var hashes []string
	db.Table("sensor_data").Order("timestamp desc").Pluck("hash", &hashes)
	respondJSON(w, http.StatusOK, hashes)
}

func GetSensorDataWrong(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondError(w, http.StatusBadRequest, "Malformed query")
}

func GetSensorDataDayHash(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	day := r.URL.Query().Get("day")
	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	var hashes []string

	db.Table("sensor_data").Where("year(timestamp) = ? AND month(timestamp) = ? AND day(timestamp) = ?", year, month, day).Order("timestamp desc").Pluck("hash", &hashes)

	respondJSON(w, http.StatusOK, hashes)

}

func GetSensorDataLast(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	sensordata_cs := model.SensorData{}
	sensordata_pv := model.SensorData{}
	sensordata_mo := model.SensorData{}
	sensordata_me := model.SensorData{}
	db.Where("sensor_id = ?", "petit-verdot").Order("timestamp desc").Limit(1).Find(&sensordata_pv)
	db.Where("sensor_id = ?", "cabernet-sauvignon").Order("timestamp desc").Limit(1).Find(&sensordata_cs)
	db.Where("sensor_id = ?", "malbec-este").Order("timestamp desc").Limit(1).Find(&sensordata_me)
	db.Where("sensor_id = ?", "malbec-oeste").Order("timestamp desc").Limit(1).Find(&sensordata_mo)

	sensordata := []model.SensorData{sensordata_cs, sensordata_pv, sensordata_mo, sensordata_me}
	respondJSON(w, http.StatusOK, sensordata)

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

	db.Select("min(timestamp) as timestamp, sensor_id,"+
		"avg(humidity2) as humidity2, avg(humidity1) as humidity1,"+
		"avg(humidity05) as humidity05, avg(humidity005) as humidity005,"+
		"max(wind_velocity) as wind_velocity, max(wind_gust) as wind_gust,"+
		"avg(wind_direction) as wind_direction, avg(pressure) as pressure,"+
		"max(rain) as rain, avg(temperature) as temperature,"+
		"avg(humidity) as humidity").Where("year(timestamp) = ? AND month(timestamp) = ?", year, month).Group("day(timestamp), sensor_id").Find(&sensordata)

	respondJSON(w, http.StatusOK, sensordata)

}

func GetSensorDataMonthHash(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	month := r.URL.Query().Get("month")
	year := r.URL.Query().Get("year")

	var hashes []string

	db.Table("sensor_data").Where("year(timestamp) = ? AND month(timestamp) = ?", year, month).Order("timestamp desc").Pluck("hash", &hashes)

	respondJSON(w, http.StatusOK, hashes)

}

func GetSensorDataYear(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	year := r.URL.Query().Get("year")

	sensordata := []model.SensorData{}

	db.Select("min(timestamp) as timestamp, sensor_id,"+
		"avg(humidity2) as humidity2, avg(humidity1) as humidity1,"+
		"avg(humidity05) as humidity05, avg(humidity005) as humidity005,"+
		"max(wind_velocity) as wind_velocity, max(wind_gust) as wind_gust,"+
		"avg(wind_direction) as wind_direction, avg(pressure) as pressure,"+
		"max(rain) as rain, avg(temperature) as temperature,"+
		"avg(humidity) as humidity").Group("month(timestamp), sensor_id").Having("year(timestamp) = ?", year).Find(&sensordata)

	respondJSON(w, http.StatusOK, sensordata)

}

func GetSensorDataYearHash(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	year := r.URL.Query().Get("year")

	var hashes []string

	db.Table("sensor_data").Where("year(timestamp) = ?", year).Pluck("hash", &hashes)

	respondJSON(w, http.StatusOK, hashes)

}
