package token

import (
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"net/http"
)

type QueryData struct {
	WinerieID string `json:"winerie_id"`
}

func GetTokensByWinerie(w http.ResponseWriter, r *http.Request) {
	var params = QueryData{}
	params.WinerieID = r.URL.Query().Get("winerie_id")

	tokens := []model.TokenWinerie{}

	query := repository.DB

	if params.WinerieID != "" {
		query = query.Where("winerie_id = ?", params.WinerieID)
	}

	err := query.Find(&tokens).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	customHTTP.ResponseJSON(w, tokens)
	return
}
