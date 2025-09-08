package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/thisusami/donation-assignment/external"
	"github.com/thisusami/donation-assignment/models"
	"github.com/thisusami/donation-assignment/service/cipher"
)

type Calculator struct {
	csvData     [][]string
	filePath    string
	omiseClient *external.OmiseClient
}

func (c *Calculator) Read() (*Calculator, error) {
	file, err := os.Open(c.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	rot128Reader, err := cipher.NewRot128Reader(file)
	if err != nil {
		return nil, fmt.Errorf("failed to create rot128 reader: %w", err)
	}
	csvReader := csv.NewReader(rot128Reader)
	csvReader.Comma = ','
	csvReader.Comment = '#'
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV data: %w", err)
	}
	c.csvData = records
	return c, nil
}
func (c *Calculator) Calculate() (*models.DonationSummary, error) {
	var (
		summary = &models.DonationSummary{}
		count   = 0
		donated []models.Donator
	)
	for i, record := range c.csvData {
		if i == 0 {
			continue
		}
		amount, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			continue
		}
		summary.Total += amount
		donated = append(donated, models.Donator{
			Name:   record[0],
			Amount: amount,
		})
		_, err = c.omiseClient.Request(record[0], record[2], record[3], record[4], record[5], amount)
		if err != nil {
			summary.Faulty += amount

		} else {
			summary.Successfully += amount
		}
		count++
	}
	summary.TopDonator = c.GetTopDonators(donated)
	if count > 0 {
		summary.Average = summary.Total / float64(count)
	}
	return summary, nil
}
func (c *Calculator) GetTopDonators(donated []models.Donator) []string {
	sort.Slice(donated, func(i, j int) bool {
		return donated[i].Amount > donated[j].Amount
	})
	var topDonators []string
	for i := 0; i < 3 && i < len(donated); i++ {
		topDonators = append(topDonators, donated[i].Name)
	}
	return topDonators
}
func NewCalculator(path string, omiseClient *external.OmiseClient) *Calculator {
	return &Calculator{
		csvData:     make([][]string, 0),
		filePath:    path,
		omiseClient: omiseClient,
	}
}
