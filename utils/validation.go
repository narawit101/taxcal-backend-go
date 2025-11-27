package utils

import (
	"errors"
	"test-backend/models"
)

func ValidateTaxRequest(req models.TaxRequest) error {
	if req.TotalIncome < 0 {
		return errors.New("totalIncome must not be negative")
	}
	if req.WHT < 0 {
		return errors.New("wht must not be negative")
	}
	if req.WHT > req.TotalIncome {
		return errors.New("wht must not exceed totalIncome")
	}
	return nil
}
