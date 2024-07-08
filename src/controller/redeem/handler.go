package redeem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/smtp"
	"strconv"

	"github.com/google/uuid"
	"github.com/jordan-wright/email"
	"github.com/openvino/openvino-api/src/config"
	customHTTP "github.com/openvino/openvino-api/src/http"
	"github.com/openvino/openvino-api/src/model"
	"github.com/openvino/openvino-api/src/repository"
	"github.com/thedevsaddam/govalidator"
)

type QueryShipping struct {
	CountryId  string `json:"country_id"`
	ProvinceId string `json:"province_id"`
	Amount     uint   `json:"amount"`
}

type CreateRedeem struct {
	PublicKey      string `json:"public_key"`
	Email          string `json:"email"`
	Name           string `json:"name"`
	Year           string `json:"year"`
	Street         string `json:"street"`
	Number         string `json:"number"`
	CountryId      string   `json:"country_id"`
	ProvinceId     string   `json:"province_id"`
	Zip            string `json:"zip"`
	TelegramId     string `json:"telegram_id"`
	Amount         uint   `json:"amount"`
	Signature      string `json:"signature"`
	BurnTxHash     string `json:"burn_tx_hash"`
	ShippingTxHash string `json:"shipping_tx_hash"`
	WinerieID      string `json:"winerie_id"`
	ShippingPaidStatus string `json:"shipping_paid_status"`
	Pickup string `json:"pickup"`
}

type QueryRedeem struct {
	Year      string `json:"year"`
	WinerieID string `json:"winerie_id"`
}

type ShippingCostResponse struct {
	Cost float64 `json:"cost"`
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
		"country_id":       []string{"required", "string"},
		"province_id":      []string{"required", "string"},
		"zip":              []string{"required", "string"},
		"telegram_id":      []string{"optional", "string"},
		"burn_tx_hash":     []string{"required", "string"},
		"shipping_tx_hash": []string{"required", "string"},
		"signature":        []string{"required", "string"},
		"winerie_id":       []string{"required", "string"},
        "shipping_paid_status": []string{"required", "string"},
		"pickup":      []string{"optional", "string"},

	}
	err := customHTTP.DecodeJSONBody(w, r, &body, rules)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var winerie model.Winerie
	err = repository.DB.First(&winerie, "id = ?", body.WinerieID).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user := model.User{
		PublicKey: body.PublicKey,
	}
	repository.DB.FirstOrCreate(&user, user)
	redeem := model.RedeemInfo{
		ID:             uuid.New().String(),
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
		WinerieID:      body.WinerieID,
		ShippingPaidStatus : body.ShippingPaidStatus,
		Watched : false,
		Status : `pending`,
		Phone : ``,
		City : ``,
		Pickup : body.Pickup,
	}
	repository.DB.Create(&redeem)

	// sender := NewGmailSender("OpenVino", config.Config.Email, config.Config.EmailPassword)

	// subject := "You have a New Redeem"
	// content := `<h1>The customer  : ` + body.Name + ` made a new redeem</h1> 
	// <a href=` + config.Config.DashboardUrl + redeem.ID + `/> Clik here to more details</a>`

	// to := []string{winerie.Email}
	// attachFiles := []string{}

	// err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	// if err != nil {
	// 	customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// Envía la notificación a Next.js
	err = sendNotification(body.Name, redeem.ID)
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
}

func GetRedeemInfo(w http.ResponseWriter, r *http.Request) {
	redeems := []model.RedeemInfo{}
	var params = QueryRedeem{}
	params.Year = r.URL.Query().Get("year")
	params.WinerieID = r.URL.Query().Get("winerie_id")

	query := repository.DB
	if params.WinerieID != "" {
		query = query.Where("winerie_id = ?", params.WinerieID)
	}
	if params.Year != "" {
		query = query.Where("year=?", params.Year)
	}
	err := query.Preload("Customer").Find(&redeems).Error
	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	customHTTP.ResponseJSON(w, redeems)
	return
}

func GetShippingCosts(w http.ResponseWriter, r *http.Request) {
	cost := model.ShippingCost{}
	var params = QueryShipping{}
	params.CountryId = r.URL.Query().Get("country_id")
	params.ProvinceId = r.URL.Query().Get("province_id")
	amount, err := strconv.ParseUint(r.URL.Query().Get("amount"), 10, 32)

	if err != nil || params.CountryId == "" || params.ProvinceId == "" {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, "The provided params are incorrect")
		return
	}

	params.Amount = uint(math.Max(float64(amount), 6))

	err = repository.DB.
		Where("country_id=? AND province_id=?",
			params.CountryId, params.ProvinceId).
		First(&cost).Error

	if err != nil {
		customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	costReturn := ShippingCostResponse{
		Cost: math.Round(((cost.BaseCost*float64(params.Amount)/6.0)+(cost.CostPerUnit*float64(params.Amount)))*100) / 100,
	}
	customHTTP.ResponseJSON(w, costReturn)
	return
}



type EmailSender interface {
	SendEmail(
		subject string,
		content string,
		to []string,
		cc []string,
		bcc []string,
		attachFiles []string,
	) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name string, fromEmailAddress string, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name:              name,
		fromEmailAddress:  fromEmailAddress,
		fromEmailPassword: fromEmailPassword,
	}
}

func (sender *GmailSender) SendEmail(
	subject string,
	content string,
	to []string,
	cc []string,
	bcc []string,
	attachFiles []string,
) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", sender.name, sender.fromEmailAddress)
	e.Subject = subject
	e.HTML = []byte(content)
	e.To = to
	e.Cc = cc
	e.Bcc = bcc

	for _, f := range attachFiles {
		_, err := e.AttachFile(f)
		if err != nil {
			return fmt.Errorf("failed to attach file %s: %w", f, err)
		}
	}

	smtpAuth := smtp.PlainAuth("", sender.fromEmailAddress, sender.fromEmailPassword, config.Config.EmailSmtp)
	return e.Send( config.Config.EmailPort, smtpAuth)
}

func sendNotification(customerName, redeemID string) error {
	url := config.Config.ServerUrl // Reemplaza con la URL correcta de tu aplicación Next.js

	data := map[string]string{
		"customerName": customerName,
		"redeemId":     redeemID,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Aquí puedes verificar la respuesta si es necesario

	return nil
}


func UpdateRedeemInfo(w http.ResponseWriter, r *http.Request) {
    if r.Body == nil {
        customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Request body is empty")
        return
    }

    var body struct {
        BurnTxHash         string `json:"burn_tx_hash"`
        ShippingTxHash     string `json:"shipping_tx_hash"`
        ShippingPaidStatus string `json:"shipping_paid_status"`
    }

    rules := govalidator.MapData{
        "burn_tx_hash":         []string{"required", "string"},
        "shipping_tx_hash":     []string{"required", "string"},
        "shipping_paid_status": []string{"required", "string"},
    }

    err := customHTTP.DecodeJSONBody(w, r, &body, rules)
    if err != nil {
        customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
        return
    }

    var redeem model.RedeemInfo
    err = repository.DB.First(&redeem, "burn_tx_hash = ?", body.BurnTxHash).Error
    if err != nil {
        customHTTP.NewErrorResponse(w, http.StatusNotFound, "Redeem not found")
        return
    }

    redeem.ShippingTxHash = body.ShippingTxHash
    redeem.ShippingPaidStatus = body.ShippingPaidStatus

    err = repository.DB.Save(&redeem).Error
    if err != nil {
        customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Failed to update redeem info")
        return
    }

    customHTTP.ResponseJSON(w, map[string]string{"message": "Redeem info updated successfully"})
}

func CreateErrorRedeemInfo(w http.ResponseWriter, r *http.Request) {
    if r.Body == nil {
        customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Request body is empty")
        return
    }

    var body struct {
        PublicKey         string `json:"public_key"`
        Email             string `json:"email"`
        Name              string `json:"name"`
        Year              string `json:"year"`
        Street            string `json:"street"`
        Number            string `json:"number"`
        CountryId         string   `json:"country_id"`
        ProvinceId        string   `json:"province_id"`
        Zip               string `json:"zip"`
        TelegramId        string `json:"telegram_id"`
        Amount            uint   `json:"amount"`
        Signature         string `json:"signature"`
        BurnTxHash        string `json:"burn_tx_hash"`
        ShippingTxHash    string `json:"shipping_tx_hash"`
        WinerieID         string `json:"winerie_id"`
        ShippingPaidStatus string `json:"shipping_paid_status"`
        ErrorMessage      string `json:"error_message"`
		Pickup 				string 	`json:"pickup"`
    }

    rules := govalidator.MapData{
        "public_key":         []string{"string"},
        "name":               []string{"string"},
        "email":              []string{"string"},
        "amount":             []string{"uint"},
        "year":               []string{"string"},
        "street":             []string{"string"},
        "number":             []string{"string"},
    	"country_id":       []string{ "string"},
		"province_id":      []string{ "string"},
        "zip":                []string{"string"},
        "telegram_id":        []string{"string"},
        "burn_tx_hash":       []string{"string"},
        "shipping_tx_hash":   []string{"string"},
        "signature":          []string{"string"},
        "winerie_id":         []string{"string"},
        "shipping_paid_status": []string{"string"},
        "error_message":      []string{"string"},
		"pickup":             []string{"string"},
    }

    err := customHTTP.DecodeJSONBody(w, r, &body, rules)
    if err != nil {
        customHTTP.NewErrorResponse(w, http.StatusBadRequest, err.Error())
        return
    }

    // Verificar si el cuerpo decodificado tiene los campos obligatorios
    if body.PublicKey == ""  {
        customHTTP.NewErrorResponse(w, http.StatusBadRequest, "Required fields are missing")
        return
    }

    user := model.User{
        PublicKey: body.PublicKey,
    }
    repository.DB.FirstOrCreate(&user, user)

    redeemLog := model.RedeemLog{
        ID:                 uuid.New().String(),
        CustomerId:         body.PublicKey,
        Customer:           user,
        Year:               body.Year,
        Street:             body.Street,
        Number:             body.Number,
        CountryId:          body.CountryId,
        ProvinceId:         body.ProvinceId,
        Zip:                body.Zip,
        TelegramId:         body.TelegramId,
        Amount:             body.Amount,
        Signature:          body.Signature,
        BurnTxHash:         body.BurnTxHash,
        ShippingTxHash:     body.ShippingTxHash,
        WinerieID:          body.WinerieID,
        ShippingPaidStatus: body.ShippingPaidStatus,
        ErrorMessage:       body.ErrorMessage,
		Pickup :            body.Pickup,
    }

    if err := repository.DB.Create(&redeemLog).Error; err != nil {
        customHTTP.NewErrorResponse(w, http.StatusInternalServerError, "Failed to create redeem log")
        return
    }

    customHTTP.ResponseJSON(w, map[string]string{"message": "Redeem log created successfully"})
}

