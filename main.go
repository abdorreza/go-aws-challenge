package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Product struct {
	Id          string `json:"Id"`
	DeviceModel string `json:"DeviceModel"`
	Name        string `json:"Name"`
	Note        string `json:"Note"`
	Serial      string `json:"Serial"`
}

func getAllData() (events.APIGatewayProxyResponse, error) {
	products := []Product{}
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

func getOneData() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-north-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	key := struct {
		ID string `dynamodbav:"id" json:"id"`
	}{ID: "A020000102"}
	avs, err := attributevalue.MarshalMap(key)
	if err != nil {
		panic(err)
	}

	out, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("devices"),
		Key:       avs,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(out.Item)
}

func main() {
	getOneData()
	fmt.Println("---------------------------------------------")
	//lambda.Start(getAllData)
	getAllData() //Nice Work

}
