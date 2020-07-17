package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/openvino/openvino-api/src/app/model"
)

type finalizeTasks struct {
	HashList []string `json:"hashes"`
}

func GetTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	tasks := []model.Task{}
	db.Find(&tasks)
	respondJSON(w, http.StatusOK, tasks)

}

func CreateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var task model.Task = model.Task{}

	if err := decoder.Decode(&task); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&task).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, task)

}

func UpdateTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var task model.Task = model.Task{}
	params := mux.Vars(r)

	db.Model(&task).Where("hash = ?", params["hash"]).Update("status", 1)
	db.Where("hash = ?", params["hash"]).First(&task)

	respondJSON(w, http.StatusOK, task)

}

func UpdateTasks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var task model.Task = model.Task{}

	var hashes finalizeTasks = finalizeTasks{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&hashes); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	db.Model(&task).Where("hash IN (?)", hashes.HashList).Update("status", 1)

	respondJSON(w, http.StatusOK, task)

}
