package db

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/abdorreza/go-aws-challenge/model"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Get All Rows
func GetAllData() (events.APIGatewayProxyResponse, error) {
	products := []model.Devices{}
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-north-1"
		return nil
	})

	if err != nil {
		panic(err)
	}
	svc := dynamodb.NewFromConfig(cfg)

	out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("devices"),
	})

	if err != nil {
		panic(err)
	}

	err = attributevalue.UnmarshalListOfMaps(out.Items, &products)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}

	productsJson, err := json.Marshal(products)
	if err != nil {
		panic(err)
	}

	resp := events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(productsJson),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	fmt.Println(resp)

	return resp, nil
}

// Get One Row
func GetOneData() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-north-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("devices"),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: "/devices/id2"},
		},
	})

	if err != nil {
		panic(err)
	}

	a, _ := json.Marshal(out.Item)
	s := string(a)

	fmt.Println(s)
}
