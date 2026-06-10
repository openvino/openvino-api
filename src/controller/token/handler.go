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

type TokenResponse struct {
	model.TokenWinerie
	Description string `json:"description"`
}

type tokenDescription struct {
	ID          string
	Description string
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

	ids := make([]string, 0, len(tokens))
	for _, token := range tokens {
		ids = append(ids, token.ID)
	}

	descriptions := []tokenDescription{}
	if len(ids) > 0 {
		err = repository.DB.Table("token_launch").
			Select("id, description").
			Where("id IN (?)", ids).
			Find(&descriptions).Error
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	descriptionByID := map[string]string{}
	for _, description := range descriptions {
		descriptionByID[description.ID] = description.Description
	}

	response := make([]TokenResponse, 0, len(tokens))
	for _, token := range tokens {
		response = append(response, TokenResponse{
			TokenWinerie: token,
			Description:  descriptionByID[token.ID],
		})
	}

	customHTTP.ResponseJSON(w, response)
	return
}
