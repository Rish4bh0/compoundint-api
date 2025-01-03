package models

type LoanRequest struct {
	InterestRate  float64 `json:"interestRate"`
	TotalAmount   float64 `json:"totalAmount"`
	DownPayment   float64 `json:"downPayment"`
	DetailedTable bool    `json:"detailedTable"`
}

type LoanSummary struct {
	TotalMonths    int     `json:"totalMonths"`
	TotalPayment   float64 `json:"totalPayment"`
	TotalPrincipal float64 `json:"totalPrincipal"`
	TotalInterest  float64 `json:"totalInterest"`
	Balance        float64 `json:"balance"`
}

type LoanResponse struct {
	MonthlyPayment float64      `json:"monthlyPayment"`
	Summary        *LoanSummary `json:"summary,omitempty"`
}
