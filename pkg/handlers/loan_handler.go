package handlers

import (
	"compoundint-api/pkg/models"
	"compoundint-api/pkg/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func LoanHandler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var loanReq models.LoanRequest
	if err := json.Unmarshal([]byte(req.Body), &loanReq); err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf("Invalid JSON: %v", err),
		}, nil
	}

	if loanReq.InterestRate <= 0 || loanReq.TotalAmount <= 0 || loanReq.DownPayment < 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid loan inputs",
		}, nil
	}

	principal := loanReq.TotalAmount - loanReq.DownPayment

	monthlyPayment := utils.MonthlyPaymentCalc(principal, loanReq.InterestRate, 30)

	var resp models.LoanResponse
	resp.MonthlyPayment = monthlyPayment

	if loanReq.DetailedTable {
		summary := utils.AmortizationDetail(principal, loanReq.InterestRate, 30, monthlyPayment)
		resp.Summary = &summary
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf("Failed to serialize response: %v", err),
		}, nil
	}

	// return client response
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(respJSON),
		Headers:    map[string]string{"Content-Type": "application/json"},
	}, nil
}
