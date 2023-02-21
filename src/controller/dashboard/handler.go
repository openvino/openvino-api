package dashboard

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
)

type QueryData struct {
	WinerieID string `json:"winerie_id"`
	Year      string `json:"year"`
}

type ResponseDashboard struct {
	Sensors  []model.SensorRecord `json:"sensor"`
	Analysis []model.AnalysisInfo `json:"analysis"`
	Task     model.Task           `json:"task"`
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {

	var params = QueryData{}
	params.WinerieID = r.URL.Query().Get("winerie_id")
	params.Year = r.URL.Query().Get("year")

	sensors := []model.SensorRecord{}
	sensordataCs := model.SensorRecord{}
	sensordataPv := model.SensorRecord{}
	sensordataMo := model.SensorRecord{}
	sensordataMe := model.SensorRecord{}
	sensorQuery := repository.DB.
		Where("winerie_id = ?", params.WinerieID).
		Where("EXTRACT(YEAR FROM timestamp) = ?", params.Year)
	sensorQuery.Where("sensor_id = ?", "petit-verdot").Order("timestamp desc").Limit(1).Find(&sensordataPv)
	sensorQuery.Where("sensor_id = ?", "cabernet-sauvignon").Order("timestamp desc").Limit(1).Find(&sensordataCs)
	sensorQuery.Where("sensor_id = ?", "malbec-este").Order("timestamp desc").Limit(1).Find(&sensordataMe)
	sensorQuery.Where("sensor_id = ?", "malbec-oeste").Order("timestamp desc").Limit(1).Find(&sensordataMo)
	sensors = []model.SensorRecord{sensordataCs, sensordataPv, sensordataMo, sensordataMe}

	task := model.Task{}
	repository.DB.
		Where("winerie_id = ?", params.WinerieID).
		Where("EXTRACT(YEAR FROM timestamp) = ?", params.Year).
		Order("end_timestamp desc").Limit(1).Find(&task)

	analysis := []model.AnalysisInfo{}
	analysisCs := model.AnalysisInfo{}
	analysisPv := model.AnalysisInfo{}
	analysisMo := model.AnalysisInfo{}
	analysisMe := model.AnalysisInfo{}

	analysisQuery := repository.DB.
		Where("winerie_id = ?", params.WinerieID).
		Where("year = ?", params.Year)
	analysisQuery.Where("grape_type = ?", "petit-verdot").Order("created_at desc").Limit(1).Find(&analysisPv)
	analysisQuery.Where("grape_type = ?", "cabernet-sauvignon").Order("created_at desc").Limit(1).Find(&analysisCs)
	analysisQuery.Where("grape_type = ?", "malbec-este").Order("created_at desc").Limit(1).Find(&analysisMe)
	analysisQuery.Where("grape_type = ?", "malbec-oeste").Order("created_at desc").Limit(1).Find(&analysisMo)
	analysis = []model.AnalysisInfo{analysisCs, analysisPv, analysisMo, analysisMe}

	dashboard := ResponseDashboard{
		Sensors:  sensors,
		Analysis: analysis,
		Task:     task,
	}
	customHTTP.ResponseJSON(w, dashboard)
	return
}
