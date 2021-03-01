package task

import (
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"github.com/thedevsaddam/govalidator"
	"log"
	"net/http"
	"time"
)

type QueryData struct {
	Harvest   string `json:"year"`
	Month     string `json:"month"`
	Day       string `json:"day"`
	PublicKey string `json:"public_key"`
}

type InsertData struct {
	Hash            string     `json:"hash"`
	PublicKey       string     `json:"public_key"`
	IniTimestamp    *time.Time `json:"ini_timestamp"`
	IniClaro        string     `json:"ini_claro"`
	IniRow          uint       `json:"ini_row"`
	IniPlant        uint       `json:"ini_plant"`
	EndTimestamp    *time.Time `json:"end_timestamp"`
	EndClaro        string     `json:"end_claro"`
	EndRow          uint       `json:"end_row"`
	EndPlant        uint       `json:"end_plant"`
	TypeId          uint       `json:"task_id"`
	ToolsUsed       []uint     `json:"tools_used"`
	Chemicals       []uint     `json:"chemicals"`
	ChemicalAmounts []float32  `json:"chemicals_amount"`
	Notes           string     `json:"notes"`
}

type ToolsData struct {
	Id uint `json:"id"`
}

func GetTasks(w http.ResponseWriter, r *http.Request) {

	var params = QueryData{}
	params.Harvest = r.URL.Query().Get("year")
	params.Month = r.URL.Query().Get("month")
	params.Day = r.URL.Query().Get("day")
	params.PublicKey = r.URL.Query().Get("public_key")

	log.Println(params)

	tasks := []model.Task{}

	query := repository.DB

	if params.Harvest != "" {
		query.Where("YEAR(ini_timestamp) = ?", params.Harvest)
	} else if params.Month != "" {
		query.Where("MONTH(ini_timestamp) = ?", params.Month)
	} else if params.Day != "" {
		query.Where("DAY(ini_timestamp) = ?", params.Day)
	} else if params.PublicKey != "" {
		query.Where("public_key = ?", params.PublicKey)
	}
	query.Preload("ToolsUsed").Preload("ChemicalsUsed").Find(&tasks)

	customHTTP.ResponseJSON(w, tasks)
	return
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	var body InsertData
	rules := govalidator.MapData{
		"hash":             []string{"required", "string"},
		"public_key":       []string{"required", "string"},
		"ini_timestamp":    []string{"required", "date"},
		"ini_claro":        []string{"required", "string"},
		"ini_row":          []string{"required", "uint"},
		"ini_plant":        []string{"required", "uint"},
		"end_timestamp":    []string{"required", "date"},
		"end_claro":        []string{"required", "string"},
		"end_row":          []string{"required", "uint"},
		"end_plant":        []string{"required", "uint"},
		"task_id":          []string{"required", "uint"},
		"tools_used":       []string{"required", "[]uint"},
		"chemicals":        []string{"required", "[]string"},
		"chemicals_amount": []string{"required", "[]float32"},
		"notes":            []string{"optional", "string"},
	}
	err := customHTTP.DecodeJSONBody(w, r, &body, rules)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	task := model.Task{
		Hash:         body.Hash,
		PublicKey:    body.PublicKey,
		IniTimestamp: body.IniTimestamp,
		IniClaro:     body.IniClaro,
		IniRow:       body.IniRow,
		IniPlant:     body.IniPlant,
		EndTimestamp: body.EndTimestamp,
		EndClaro:     body.EndClaro,
		EndRow:       body.EndRow,
		EndPlant:     body.EndPlant,
		TypeId:       body.TypeId,
		Notes:        body.Notes,
	}

	for _, element := range body.ToolsUsed {
		task.ToolsUsed = append(task.ToolsUsed, model.Tools{
			Id:       element,
			TaskHash: body.Hash,
		})
	}

	for i, element := range body.Chemicals {
		task.ChemicalsUsed = append(task.ChemicalsUsed, model.Chemicals{
			Id:       element,
			Amount:   body.ChemicalAmounts[i],
			TaskHash: body.Hash,
		})
	}

	log.Println(task)

	repository.DB.Create(task)

}
