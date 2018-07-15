package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Received body: ", request.Body)

	var c = new(contact)

	err := json.Unmarshal([]byte(request.Body), c)
	if err != nil {
		log.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       "Invalid Contact",
			StatusCode: 422,
		}, nil
	}

	service := contactService{contact: c}
	_, err = service.create()

	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 422,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Hello " + request.Body,
		StatusCode: 200,
	}, nil
}
