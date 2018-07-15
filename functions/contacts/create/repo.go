package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws/session"
)

// contactRepo - collection of contacts
type contactRepo struct{}

// contactDAO - database access management for contacts
var contactDAO = contactRepo{}

// contactTableName - where contacts are stored
var contactTableName = "contacts"

func (c *contactRepo) session() (*dynamodb.DynamoDB, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		log.Println(err)
	}

	return dynamodb.New(session), err
}

// Insert - persists contact into storage
func (c *contactRepo) insert(contact *contact) (*contact, error) {
	item, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		log.Println(err)
		return contact, err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(contactTableName),
	}

	db, err := c.session()
	if err != nil {
		log.Println(err)
		return contact, err
	}

	_, err = db.PutItem(input)
	if err != nil {
		log.Println(err)
		return contact, err
	}

	return contact, err
}
