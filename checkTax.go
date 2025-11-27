package services

func CheckTax(income float64) {
	if income <= 0 {
		return
	}

	taxPercent := 0.0
	if income > 2000000 {
		taxPercent = 35
	} else if income < 2000001 && income >= 1000001 {
		taxPercent = 20
	} else if income < 1000001 && income >= 500001 {
		taxPercent = 15
	} else if income < 500001 && income >= 150001 {
		taxPercent = 10
	} else {
		taxPercent = 0
	}
}
