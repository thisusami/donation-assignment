package models

type Request struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
	Card     Card   `json:"card"`
}
type Card struct {
	Number          string `json:"number"`
	ExpirationMonth int    `json:"expiration_month"`
	ExpirationYear  int    `json:"expiration_year"`
	SecurityCode    string `json:"security_code"`
	Name            string `json:"name"`
}
type Response struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
	Status string `json:"status"`
}
