package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
)

func GetHello(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"hello": "world"})
}

func CreateHello(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"hello": "world"})
}

func UpdateHello(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"hello": "world"})
}

func DeleteHello(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, map[string]string{"hello": "world"})
}
