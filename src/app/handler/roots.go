package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/openvino/openvino-api/src/app/model"
)

func GetRoot(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	root := mux.Vars(r)["root"]
	tx_hash := []model.Root{}

	db.Where("root = ?", root).Find(&tx_hash)
	respondJSON(w, http.StatusOK, tx_hash)
	
}