package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

const (
	// PkName - SkName - GsiName - LsiName - TableName should edit these field as you use
	PkName       = "PK"
	SkName       = "SK"
	GsiName      = "GSI"
	LsiName      = "LSI"
	GsiIndexName = "GlobalSecondaryIndex"
	LsiIndexName = "LocalSecondaryIndex"
	TableName    = "MyTable"

	// GlobalWriteReadCap should be equal or higher than PartitionWriteReadCap
	GlobalWriteReadCap    = 10
	PartitionWriteReadCap = 10

	Location = "eu-west-2" // your location!!
)

func main() {

	cfg, err := config.LoadDefaultConfig(context.TODO(), func(opts *config.LoadOptions) error {
		opts.Region = Location
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)
	out, err := svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName: aws.String(TableName),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String(PkName),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String(SkName),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String(PkName),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String(SkName),
				KeyType:       types.KeyTypeRange,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(PartitionWriteReadCap),
			WriteCapacityUnits: aws.Int64(PartitionWriteReadCap),
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}

/*func main() {
	lambda.Start(LambdaHandler)
}*/
