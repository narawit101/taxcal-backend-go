package services

type CalculationResult struct {
	TotalIncome float64
	WHT         float64
}

func Calculations(totalIncome float64, wht float64) CalculationResult {

	result := CalculationResult{
		TotalIncome: totalIncome,
		WHT:         wht,
	}

	return result
}
