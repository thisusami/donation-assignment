package models

type DonationSummary struct {
	Total        float64
	Successfully float64
	Faulty       float64
	Average      float64
	TopDonator   []string
}
type Donator struct {
	Name   string
	Amount float64
}
