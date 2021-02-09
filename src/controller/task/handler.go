package task

import (
	"github.com/thedevsaddam/govalidator"
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"log"
	"net/http"
	"time"
)

type QueryData struct {
	Harvest   string  `json:"year"`
	Month     string  `json:"month"`
	Day		  string  `json:"day"`
}

type InsertData struct {
	Hash			string   	`json:"hash"`
	IniTimestamp 	*time.Time 	`json:"ini_timestamp"`
	IniClaro	 	string		`json:"ini_claro"`
	IniRow	 	 	uint		`json:"ini_row"`
	IniPlant	 	uint		`json:"ini_plant"`
	EndTimestamp 	*time.Time 	`json:"end_timestamp"`
	EndClaro	 	string		`json:"end_claro"`
	EndRow	 	 	uint		`json:"end_row"`
	EndPlant	 	uint		`json:"end_plant"`
	TypeId	  		uint 		`json:"task_id"`
	ToolsUsed	  	[]uint		`json:"tools_used"`
	Chemicals 		[]uint		`json:"chemicals"`
	ChemicalAmounts []float32	`json:"chemicals_amount"`
	Notes			string		`json:"notes"`
}

type ToolsData struct {
	Id				uint		`json:"id"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	var params = QueryData{}
	params.Harvest = r.URL.Query().Get("year")
	params.Month = r.URL.Query().Get("month")
	params.Day = r.URL.Query().Get("day")

	log.Println(params)

	tasks := []model.Task{}

	if params.Day == "" && params.Month == "" && params.Harvest != "" {
		repository.DB.
			Where("YEAR(ini_timestamp) = ?", params.Harvest).
			Preload("ToolsUsed").Preload("ChemicalsUsed").
			Find(&tasks)
	} else if params.Day == "" && params.Month != "" && params.Harvest != "" {
		repository.DB.
			Preload("ToolsUsed").Preload("ChemicalsUsed").
			Where("MONTH(ini_timestamp) = ? AND YEAR(ini_timestamp) = ?", params.Month, params.Harvest).
			Find(&tasks);
	} else if params.Day != "" && params.Month != "" && params.Harvest != "" {
		repository.DB.
			Preload("ToolsUsed").Preload("ChemicalsUsed").
			Where("DAY(ini_timestamp) = ? AND MONTH(ini_timestamp) = ? AND YEAR(ini_timestamp) = ?", params.Day, params.Month, params.Harvest).
			Find(&tasks);
	} else {
		repository.DB.
			Preload("ToolsUsed").Preload("ChemicalsUsed").
			Find(&tasks);
	}
	customHTTP.ResponseJSON(w, tasks)
	return
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var body InsertData
	rules := govalidator.MapData {
		"hash": []string{"required", "string"},
		"ini_timestamp": []string {"required", "date"},
		"ini_claro": []string{"required", "string"},
		"ini_row": []string{"required", "uint"},
		"ini_plant": []string{"required", "uint"},
		"end_timestamp": []string {"required", "date"},
		"end_claro": []string{"required", "string"},
		"end_row": []string{"required", "uint"},
		"end_plant": []string{"required", "uint"},
		"task_id": []string{"required", "uint"},
		"tools_used": []string{"required", "[]uint"},
		"chemicals": []string{"required", "[]string"},
		"chemicals_amount": []string{"required", "[]float32"},
		"notes": []string{"optional", "string"},
	}
	err := customHTTP.DecodeJSONBody(w, r, &body, rules)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Wrong query")
		return
	}
	task := model.Task{
		Hash: body.Hash,
		IniTimestamp: body.IniTimestamp,
		IniClaro: body.IniClaro,
		IniRow: body.IniRow,
		IniPlant: body.IniPlant,
		EndTimestamp: body.EndTimestamp,
		EndClaro: body.EndClaro,
		EndRow: body.EndRow,
		EndPlant: body.EndPlant,
		TypeId: body.TypeId,
		Notes: body.Notes,
	}

	for _, element := range body.ToolsUsed {
		task.ToolsUsed = append(task.ToolsUsed, model.Tools{
			Id: element,
			TaskHash: body.Hash,
		})
	}

	for i, element := range body.Chemicals {
		task.ChemicalsUsed = append(task.ChemicalsUsed, model.Chemicals{
			Id: element,
			Amount: body.ChemicalAmounts[i],
			TaskHash: body.Hash,
		})
	}

	log.Println(task)

	repository.DB.Create(task)

}
