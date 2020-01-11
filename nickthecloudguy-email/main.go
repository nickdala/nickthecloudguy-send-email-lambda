package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type BodyRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Response struct {
	Message string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))
	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	bodyRequest := BodyRequest{}

	// Unmarshal the json
	if err := json.Unmarshal([]byte(request.Body), &bodyRequest); err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	if err := sendEmail(bodyRequest.Name, bodyRequest.Email, bodyRequest.Message); err != nil {
		//return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, nil
		return events.APIGatewayProxyResponse{}, err
	}

	response := &Response{
		Message: fmt.Sprintf("Sent email to %v with email %v", bodyRequest.Name, bodyRequest.Email),
	}

	responseBody, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "https://www.nickthecloudguy.com",
		},
		Body:       string(responseBody),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(handler)
}
