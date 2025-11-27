package services

import (
	"test-backend/models"
)

func CalculateTax(req models.TaxRequest) models.TaxResponse {
	// STEP 1: Base Allowance
	baseAllowance := 60000.0

	// STEP 2: Allowance: Donation (max 100,000)
	donation := 0.0
	for _, a := range req.Allowances {
		if a.AllowanceType == "donation" {
			donation = a.Amount
		}
	}

	totalAllowance := baseAllowance + donation

	taxable := req.TotalIncome - totalAllowance
	if taxable < 0 {
		taxable = 0
	}

	brackets := []struct {
		Limit float64
		Rate  float64
		Label string
	}{
		{150000, 0.00, "0-150,000"},
		{500000, 0.10, "150,001-500,000"},
		{1000000, 0.15, "500,001-1,000,000"},
		{2000000, 0.20, "1,000,001-2,000,000"},
		{1e18, 0.35, "2,000,001 ขึ้นไป"},
	}

	var tax float64
	remaining := taxable
	previous := 0.0

	taxLevels := []models.TaxLevel{}

	for _, b := range brackets {
		if remaining <= 0 {
			taxLevels = append(taxLevels, models.TaxLevel{Level: b.Label, Tax: 0})
			continue
		}

		maxAmount := b.Limit - previous
		if maxAmount < 0 {
			maxAmount = 0
		}

		taxedAmount := 0.0
		if remaining > maxAmount {
			taxedAmount = maxAmount * b.Rate
			tax += taxedAmount
			remaining -= maxAmount
		} else {
			taxedAmount = remaining * b.Rate
			tax += taxedAmount
			remaining = 0
		}

		taxLevels = append(taxLevels, models.TaxLevel{
			Level: b.Label,
			Tax:   taxedAmount,
		})

		previous = b.Limit
	}

	// STEP 3: WHT
	finalTax := tax - req.WHT

	return models.TaxResponse{
		Tax:      finalTax,
		TaxLevel: taxLevels,
	}
}
