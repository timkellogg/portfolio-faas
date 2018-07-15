package main

import (
	"github.com/aws/aws-lambda-go/events"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}
