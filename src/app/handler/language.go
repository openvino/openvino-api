package handler

import (
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
	"io/ioutil"
	"encoding/json"
	"os"
)

func GetLanguage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	lang := mux.Vars(r)["lang"]
	jsonFile, err := os.Open("/lang/" + lang + ".json" )

	if err != nil {
       respondError(w, http.StatusBadRequest, err.Error())
       return
    }

    defer jsonFile.Close()
    byteValue, _ := ioutil.ReadAll(jsonFile)
    var f interface{}
    json.Unmarshal(byteValue, &f)

	respondJSON(w, http.StatusOK, f)

}