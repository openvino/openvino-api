package language

import (
	"encoding/json"
	"github.com/gorilla/mux"
	customHTTP "github.com/openvino/openvino-api/src/http"
	"io/ioutil"
	"net/http"
	"os"
)

func GetLanguage(w http.ResponseWriter, r *http.Request) {
	lang := mux.Vars(r)["lang"]
	jsonFile, err := os.Open("/languages/" + lang + ".json" )
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var f interface{}
	json.Unmarshal(byteValue, &f)
	customHTTP.ResponseJSON(w, f)
	return
}

