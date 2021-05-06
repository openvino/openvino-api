package redeem

import (
	"net/http"

	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"github.com/thedevsaddam/govalidator"
)

type QueryShipping struct {
	CountryId  string `json:"country_id"`
	ProvinceId string `json:"province_id"`
	Amount     string `json:"amount"`
}

type CreateRedeem struct {
	PublicKey      string `json:"public_key"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Year           string `json:"year"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	CountryId      uint   `json:"country_id"`
	ProvinceId     uint   `json:"province_id"`
	Zip            string `json:"zip"`
	TelegramId     string `json:"telegram_id"`
	Amount         uint   `json:"amount"`
	Signature      string `json:"signature"`
	BurnTxHash     string `json:"burn_tx_hash"`
	ShippingTxHash string `json:"shipping_tx_hash"`
}

type QueryRedeem struct {
	Year string `json:"year"`
}

func CreateReedemInfo(w http.ResponseWriter, r *http.Request) {
	var body CreateRedeem
	rules := govalidator.MapData{
		"public_key":       []string{"required", "string"},
		"name":             []string{"required", "string"},
		"email":            []string{"required", "string"},
		"amount":           []string{"required", "uint"},
		"year":             []string{"required", "string"},
		"street":           []string{"required", "string"},
		"number":           []string{"required", "string"},
		"country_id":       []string{"required", "uint"},
		"province_id":      []string{"required", "uint"},
		"zip":              []string{"required", "string"},
		"telegram_id":      []string{"optional", "string"},
		"burn_tx_hash":     []string{"required", "string"},
		"shipping_tx_hash": []string{"required", "string"},
		"signature":        []string{"required", "string"},
	}
	err := customHTTP.DecodeJSONBody(w, r, &body, rules)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	user := model.User{
		PublicKey: body.PublicKey,
		Name:      body.Name,
		Email:     body.Email,
	}
	repository.DB.FirstOrCreate(&user, user)
	redeem := model.RedeemInfo{
		CustomerId:     body.PublicKey,
		Customer:       user,
		Year:           body.Year,
		Street:         body.Street,
		Number:         body.Number,
		CountryId:      body.CountryId,
		ProvinceId:     body.ProvinceId,
		Zip:            body.Zip,
		TelegramId:     body.TelegramId,
		Amount:         body.Amount,
		Signature:      body.Signature,
		BurnTxHash:     body.BurnTxHash,
		ShippingTxHash: body.ShippingTxHash,
	}
	repository.DB.Create(&redeem)
}

func GetRedeemInfo(w http.ResponseWriter, r *http.Request) {
	redeems := []model.RedeemInfo{}
	var params = QueryRedeem{}
	params.Year = r.URL.Query().Get("year")
	if params.Year != "" {
		err := repository.DB.
			Where("year=?", params.Year).
			Preload("Customer").
			Find(&redeems).Error
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		err := repository.DB.Preload("Customer").Find(&redeems).Error
		if err != nil {
			customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}
	customHTTP.ResponseJSON(w, redeems)
	return
}

func GetShippingCosts(w http.ResponseWriter, r *http.Request) {
	cost := model.ShippingCost{}
	var params = QueryShipping{}
	params.CountryId = r.URL.Query().Get("country_id")
	params.ProvinceId = r.URL.Query().Get("province_id")
	params.Amount = r.URL.Query().Get("amount")
	err := repository.DB.
		Where("country_id=? AND province_id=? AND amount >= ?",
			params.CountryId, params.ProvinceId, params.Amount).
		Order("amount asc").
		First(&cost).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	customHTTP.ResponseJSON(w, cost)
	return
}
