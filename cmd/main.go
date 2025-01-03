package main

import (
	"compoundint-api/pkg/handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Entry point
	lambda.Start(handlers.LoanHandler)
}
