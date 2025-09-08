package service

import "github.com/thisusami/donation-assignment/models"

type CalculateInterface interface {
	Calculate() (*models.DonationSummary, error)
	Read() (*Calculator, error)
}

type SummaryInterface interface {
	Summarize() error
}
