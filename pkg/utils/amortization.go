package utils

import (
	"compoundint-api/pkg/models"
	"math"
)

// monthly payment calculation
func MonthlyPaymentCalc(principal, rate, term float64) float64 {
	monthlyRate := rate / 12 / 100
	months := term * 12
	return principal * (monthlyRate * math.Pow(1+monthlyRate, months)) / (math.Pow(1+monthlyRate, months) - 1)
}

// generating detailed amortization
func AmortizationDetail(principal, rate, term, monthlyPayment float64) models.LoanSummary {
	balance := principal
	totalPayment := 0.0
	totalPrincipal := 0.0
	totalInterest := 0.0
	totalMonths := int(term * 12)

	for month := 1; month <= totalMonths; month++ {
		interest := balance * (rate / 12 / 100)
		principalPayment := monthlyPayment - interest
		balance -= principalPayment

		if balance < 0 {
			balance = 0
		}

		totalPayment += monthlyPayment
		totalPrincipal += principalPayment
		totalInterest += interest
	}

	return models.LoanSummary{
		TotalMonths:    totalMonths,
		TotalPayment:   totalPayment,
		TotalPrincipal: totalPrincipal,
		TotalInterest:  totalInterest,
		Balance:        balance,
	}
}
