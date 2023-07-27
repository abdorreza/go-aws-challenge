package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Database struct {
	client    *dynamodb.DynamoDB
	tablename string
}

type Movie struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// Add other attributes as needed
}

func getFromDynamoDB(ctx context.Context, key string) (*Item, error) {
	// Create a new AWS session
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("YOUR_REGION"), // Replace with your desired AWS region
	})
	if err != nil {
		return nil, err
	}

	// Create a new DynamoDB client
	svc := dynamodb.New(sess)

	// Define the input parameters for GetItem operation
	params := &dynamodb.GetItemInput{
		TableName: aws.String("YOUR_TABLE_NAME"), // Replace with your DynamoDB table name
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key),
			},
		},
	}

	// Execute the GetItem operation
	result, err := svc.GetItem(params)
	if err != nil {
		return nil, err
	}

	// Unmarshal the DynamoDB attribute values into the Item struct
	item := Item{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func LambdaHandler(ctx context.Context, request interface{}) (*Item, error) {
	// Type assertion to extract the key from the request
	key, ok := request.(string)
	if !ok {
		return nil, fmt.Errorf("invalid request format")
	}

	// Call the function to get the item from DynamoDB
	item, err := getFromDynamoDB(ctx, key)
	if err != nil {
		return nil, err
	}

	return item, nil
}

/*func main() {
	lambda.Start(LambdaHandler)
}*/
