package main

import (
	"bytes"
	"encoding/json"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"pack-sales-calculator/calculator"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	var buf bytes.Buffer

	defaultPacks := []int{250,500,1000,2000,5000}
	orderString := request.QueryStringParameters["order"]

	order, err := strconv.Atoi(orderString)
	if err != nil {
		return Response{StatusCode: 500}, err
	}

	packsRequired := calculator.PackCalculator(order, defaultPacks)

	body, err := json.Marshal(packsRequired)
	if err != nil {
		return Response{StatusCode: 500}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
