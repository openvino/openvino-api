package winerie

import (
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type WinerieRequest struct {
	Name         string `json:"name"`
	Website      string `json:"website"`
	Image        string `json:"image"`
	PrimaryColor string `json:"primary_color"`
}

type WinerieResponse struct {
	Name         string `json:"name"`
	Website      string `json:"website"`
	Image        string `json:"image"`
	PrimaryColor string `json:"primary_color"`
}

func CreateWinerie(w http.ResponseWriter, r *http.Request) {
	var body WinerieRequest
	rules := govalidator.MapData{
		"name":          []string{"required", "string"},
		"website":       []string{"required", "string"},
		"image":         []string{"required", "string"},
		"primary_color": []string{"required", "string"},
	}
	err := customHTTP.DecodeJSONBody(w, r, &body, rules)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	winerie := model.Winerie{
		Name:         body.Name,
		Website:      body.Website,
		Image:        body.Image,
		PrimaryColor: body.PrimaryColor,
	}
	repository.DB.Create(&winerie)
}

func GetWineries(w http.ResponseWriter, r *http.Request) {
	wineries := []model.Winerie{}
	err := repository.DB.Find(&wineries).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	customHTTP.ResponseJSON(w, wineries)
	return
}
