package external

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/thisusami/donation-assignment/models"
)

type OmiseClient struct {
	SecretKey string
	BaseURL   string
	Client    *http.Client
}

func (o *OmiseClient) Request(name, ccNumber, cvv, expMonth, expYear string, amount float64) (models.Response, error) {
	amountSatang := int64(amount * 100)
	var chargeResp models.Response
	expMonthInt, _ := strconv.Atoi(expMonth)
	expYearInt, _ := strconv.Atoi(expYear)
	card := models.Card{
		Number:          ccNumber,
		ExpirationMonth: expMonthInt,
		ExpirationYear:  expYearInt,
		SecurityCode:    cvv,
		Name:            name,
	}

	reqData := models.Request{
		Amount:   amountSatang,
		Currency: "thb",
		Card:     card,
	}
	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return chargeResp, fmt.Errorf("failed to marshal request data: %w", err)
	}
	req, err := http.NewRequest("POST", o.BaseURL+"/charges", bytes.NewBuffer(jsonData))
	if err != nil {
		return chargeResp, fmt.Errorf("failed to create request: %w", err)
	}
	auth := base64.StdEncoding.EncodeToString([]byte(o.SecretKey + ":"))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")
	resp, err := o.Client.Do(req)
	if err != nil {
		return chargeResp, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&chargeResp); err != nil {
		return chargeResp, fmt.Errorf("failed to decode response: %w", err)
	}

	return chargeResp, nil
}
func NewOmiseClient() *OmiseClient {

	return &OmiseClient{
		SecretKey: os.Getenv("OMISE_SECRET_KEY"),
		BaseURL:   "https://api.omise.co",
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
