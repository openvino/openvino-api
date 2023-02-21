package winerie

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"

	"github.com/google/uuid"
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"github.com/thedevsaddam/govalidator"
)

type WinerieRequest struct {
	ID           string `json:"id"`
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
	Secret       string `json:"secret"`
}

func CreateWinerie(w http.ResponseWriter, r *http.Request) {
	var body WinerieRequest
	rules := govalidator.MapData{
		"id":            []string{"required", "string"},
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

	secret, err := uuid.NewUUID()
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	secretHash := sha256.Sum256([]byte(secret.String()))

	winerie := model.Winerie{
		ID:           body.ID,
		Name:         body.Name,
		Website:      body.Website,
		Image:        body.Image,
		PrimaryColor: body.PrimaryColor,
		Secret:       hex.EncodeToString(secretHash[:]),
	}
	repository.DB.Create(&winerie)

	winerie.Secret = secret.String()
	customHTTP.ResponseJSON(w, winerie)
	return
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
