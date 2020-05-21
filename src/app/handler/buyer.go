package handler

import (
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/openvino/openvino-api/src/app/model"
)

func CreateBuyer(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	var buyer model.Buyer = model.Buyer{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&buyer); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&buyer).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	respondJSON(w, http.StatusOK, buyer)
	
}

func GetBuyers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	
	buyers := []model.Buyer{}

	db.Find(&buyers)
	respondJSON(w, http.StatusOK, buyers)
	
}