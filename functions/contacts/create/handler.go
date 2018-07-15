package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Received body: ", request.Body)

	var c contact

	err := json.Unmarshal([]byte(request.Body), c)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       "Invalid Contact",
			StatusCode: 422,
		}, nil
	}

	service := contactService{contact: &c}
	service.create()

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}
