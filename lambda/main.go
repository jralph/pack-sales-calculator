package main

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"

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

type PackOrder struct {
	Size int `json:size`
	Amount int `json:amount`
}

type CalculateResponse struct {
	Packs []PackOrder `json:packs`
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	var buf bytes.Buffer

	defaultPacks := []int{250,500,1000,2000,5000}
	orderString := request.QueryStringParameters["order"]

	order, err := strconv.Atoi(orderString)
	if err != nil {
		resp := map[string]string{
			"message": "please specify a valid number for the order request parameter",
		}
		body, err := json.Marshal(resp)
		return Response{StatusCode: 400, Body: string(body)}, err
	}

	if order > 999999999999 {
		resp := map[string]string{
			"message": "web requests are limited to order sizes of 999999999999 for performance and cost",
		}
		body, err := json.Marshal(resp)
		return Response{StatusCode: 400, Body: string(body)}, err
	}

	customPacks := request.QueryStringParameters["packs"]

	if len(customPacks) > 0 {
		packStrings := strings.Split(customPacks, ",")

		var packs []int
		for _, pack := range packStrings {
			size, err := strconv.Atoi(pack)
			if err != nil {
				resp := map[string]string{
					"message": "please specify a comma seperated list of numbers for the packs query parameter",
				}
				body, err := json.Marshal(resp)
				return Response{StatusCode: 400, Body: string(body)}, err
			}

			packs = append(packs, size)
		}
		defaultPacks = packs
	}

	if order < 1 {
		resp := map[string]string{
			"message": "please specify a number greater than 0 for the order request parameter",
		}
		body, err := json.Marshal(resp)
		return Response{StatusCode: 400, Body: string(body)}, err
	}

	packsRequired := calculator.PackCalculator(order, defaultPacks)

	respData := CalculateResponse{
		Packs: []PackOrder{},
	}

	for pack, amount := range packsRequired {
		respData.Packs = append(respData.Packs, PackOrder{
			Size: pack,
			Amount: amount,
		})
	}

	body, err := json.Marshal(respData)
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
