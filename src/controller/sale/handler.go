package sale

import (
	"github.com/thedevsaddam/govalidator"
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"net/http"
)

type SaleRequest struct {
	PublicKey string 	`json:"public_key"`
	Name string			`json:"name"`
	Email string		`json:"email"`
	Amount int			`json:"amount"`
}

type SaleResponse struct {
	PublicKey string 	`json:"public_key"`
	Name string			`json:"name"`
	Email string		`json:"email"`
	Amount int			`json:"amount"`
}

func CreateSale(w http.ResponseWriter, r *http.Request) {
	var body SaleRequest
	rules := govalidator.MapData {
		"public_key": []string{"required", "string"},
		"name": []string {"optional", "string"},
		"email": []string{"required", "string"},
		"amount": []string{"required", "int"},
	}
	err := customHTTP.DecodeJSONBody(w, r, &body, rules)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	user := model.User{
		PublicKey: body.PublicKey,
		Name: body.Name,
		Email: body.Email,
	}
	repository.DB.FirstOrCreate(&user, user)
	sale := model.Sale {
		CustomerId: body.PublicKey,
		Customer: user,
		Amount: body.Amount,
	}
	repository.DB.Create(&sale)
}

func GetSales(w http.ResponseWriter, r *http.Request) {
	sales := []model.Sale{}
	err := repository.DB.Preload("Customer").Find(&sales).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	var sales_response []SaleResponse
	for _, element := range sales {
		sales_response = append(sales_response, SaleResponse{
			PublicKey: element.CustomerId,
			Name: element.Customer.Name,
			Email: element.Customer.Email,
			Amount: element.Amount,
		})
	}
	customHTTP.ResponseJSON(w, sales_response)
	return
}

