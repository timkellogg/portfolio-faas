package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
import "github.com/guregu/dynamo"

	"github.com/aws/aws-sdk-go/aws/session"
)

// contactRepo - collection of contacts
type contactRepo struct{}

// contactDAO - database access management for contacts
var contactDAO = contactRepo{}

// contactTableName - where contacts are stored
var contactTableName = "contacts"

func (c *contactRepo) session() (*dynamodb.DynamoDB, error) {
	// session, err := session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-1"),
	// })
	// if err != nil {
	// 	log.Println(err)
	// }

	// return dynamodb.New(session), err

	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String("us-east-1"),
	})
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

func (c *contactRepo) findBy(contact *contact, field string, value string) (*contact, error) {
	// filter := expression.Name(field).Equal(expression.Value(value))
	// projection := expression.NamesList(expression.Name("id"))

	// express, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
	// if err != nil {
	// 	log.Println(err)
	// 	return contact, err
	// }

	// params := {

	// }

	// db.GetItem(&dynamodb.GetItemInput{
	// 	TableName:       aws.String(contactTableName),
	// 	AttributesToGet: []*string{aws.String(field)},

	// })

	db, err := c.session()
	if err != nil {
		log.Println(err)
		return contact, err
	}

	db.Table
	// result, err := db.GetItem(&dynamodb.GetItemInput{
	// 	TableName: aws.String(contactTableName),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		field: {N: aws.String(value)},
	// 	},
	// })

	// err = dynamodbattribute.UnmarshalMap(result.Item, &contact)
	// if err != nil {

	// }
}
