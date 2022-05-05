package dashboard

import (
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"net/http"
)

type QueryData struct {
	WinerieID string `json:"winerie_id"`
}

type ResponseDashboard struct {
	Sensors  []model.SensorRecord `json:"sensor"`
	Analysis []model.AnalysisInfo `json:"analysis"`
	Task     model.Task           `json:"task"`
}

func GetDashboard(w http.ResponseWriter, r *http.Request) {

	var params = QueryData{}
	params.WinerieID = r.URL.Query().Get("winerie_id")

	stm := repository.DB
	if params.WinerieID != "" {
		stm = stm.Where("winerie_id = ?", params.WinerieID)
	}

	sensors := []model.SensorRecord{}
	sensordataCs := model.SensorRecord{}
	sensordataPv := model.SensorRecord{}
	sensordataMo := model.SensorRecord{}
	sensordataMe := model.SensorRecord{}
	stm.Where("sensor_id = ?", "petit-verdot").Order("timestamp desc").Limit(1).Find(&sensordataPv)
	stm.Where("sensor_id = ?", "cabernet-sauvignon").Order("timestamp desc").Limit(1).Find(&sensordataCs)
	stm.Where("sensor_id = ?", "malbec-este").Order("timestamp desc").Limit(1).Find(&sensordataMe)
	stm.Where("sensor_id = ?", "malbec-oeste").Order("timestamp desc").Limit(1).Find(&sensordataMo)
	sensors = []model.SensorRecord{sensordataCs, sensordataPv, sensordataMo, sensordataMe}

	task := model.Task{}
	stm.Order("end_timestamp desc").Limit(1).Find(&task)

	analysis := []model.AnalysisInfo{}
	analysisCs := model.AnalysisInfo{}
	analysisPv := model.AnalysisInfo{}
	analysisMo := model.AnalysisInfo{}
	analysisMe := model.AnalysisInfo{}
	stm.Where("grape_type = ?", "petit-verdot").Order("created_at desc").Limit(1).Find(&analysisPv)
	stm.Where("grape_type = ?", "cabernet-sauvignon").Order("created_at desc").Limit(1).Find(&analysisCs)
	stm.Where("grape_type = ?", "malbec-este").Order("created_at desc").Limit(1).Find(&analysisMe)
	stm.Where("grape_type = ?", "malbec-oeste").Order("created_at desc").Limit(1).Find(&analysisMo)
	analysis = []model.AnalysisInfo{analysisCs, analysisPv, analysisMo, analysisMe}

	dashboard := ResponseDashboard{
		Sensors:  sensors,
		Analysis: analysis,
		Task:     task,
	}
	customHTTP.ResponseJSON(w, dashboard)
	return
}
