package expense

import (
	"net/http"
	"time"

	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"github.com/thedevsaddam/govalidator"
)

type QueryData struct {
	Token     string `json:"token_id"`
	Category  string `json:"category_id"`
	WinerieID string `json:"winerie_id"`
}

type InsertData struct {
	Token       uint       `json:"token_id"`
	Timestamp   *time.Time `json:"timestamp"`
	TypeId      uint       `json:"expense_id"`
	Description string     `json:"description"`
	Value       float32    `json:"value"`
	WinerieID   string     `json:"winerie_id"`
}

type Sums struct {
	Value uint `json:"value"`
}

type ReturnExpenses struct {
	Expenses        []model.Expense `json:"expenses"`
	TotalTokensYear uint            `json:"total_tokens_year"`
	TotalTokens     uint            `json:"total_tokens"`
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {

	var params = QueryData{}
	params.Token = r.URL.Query().Get("token_id")
	params.Category = r.URL.Query().Get("category_id")
	params.WinerieID = r.URL.Query().Get("winerie_id")

	if params.Token == "" {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "The query has to specify at least a token_id")
	}

	var expenses []model.Expense

	query := repository.DB

	if params.Token != "" {
		query = query.Where("token = ? OR token = 0", params.Token)
	}

	if params.Category != "" {
		query = query.Where("type_id = ?", params.Category)
	}

	if params.WinerieID != "" {
		query = query.Where("winerie_id = ?", params.WinerieID)
	}

	query.Order("timestamp desc").Find(&expenses)

	var TokenAmount Sums
	var TotalTokens Sums

	repository.DB.Table("tokens").Select("amount as value").Where("id=?", params.Token).First(&TokenAmount)
	repository.DB.Table("tokens").Select("SUM(amount) as value").First(&TotalTokens)

	customHTTP.ResponseJSON(w, ReturnExpenses{
		Expenses:        expenses,
		TotalTokensYear: TokenAmount.Value,
		TotalTokens:     TotalTokens.Value,
	})

	return
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {

	var body InsertData
	rules := govalidator.MapData{
		"token_id":    []string{"required", "uint"},
		"timestamp":   []string{"required", "date"},
		"expense_id":  []string{"required", "uint"},
		"description": []string{"required", "string"},
		"value":       []string{"required", "float32"},
		"winerie_id":  []string{"required", "int"},
	}

	err := customHTTP.DecodeJSONBody(w, r, &body, rules)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var winerie model.Winerie
	err = repository.DB.First(&winerie, body.WinerieID).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	expense := model.Expense{
		Hash:        "",
		Token:       body.Token,
		Timestamp:   body.Timestamp,
		TypeId:      body.TypeId,
		Description: body.Description,
		Value:       body.Value,
		WinerieID:   body.WinerieID,
	}

	repository.DB.Create(expense)
	customHTTP.ResponseJSON(w, expense)
	return

}
