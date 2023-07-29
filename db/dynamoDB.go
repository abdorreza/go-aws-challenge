package db

import (
	"context"
	"sync"

	"github.com/abdorreza/go-aws-challenge/config"
	"github.com/abdorreza/go-aws-challenge/model"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

var once sync.Once
var dynamodbClient dynamodbHandler

// Return in Sigleton Manner, Return Dynaodb Client
func loadClient(ctx context.Context) error {
	if dynamodbClient != nil {
		return nil
	}
	var err error
	once.Do(func() {
		var cfg aws.Config
		cfg, err = awsConfig.LoadDefaultConfig(ctx, func(o *awsConfig.LoadOptions) error {
			o.Region = config.AWSRegion
			return nil
		})
		if err != nil {
			return
		}

		dynamodbClient = dynamodb.NewFromConfig(cfg)

	})
	if err != nil {
		return err
	}
	return nil
}

// TODO
// 1) define struct

type myStruct struct {
}

func NewMyStruct() (myStruct, error) {
	err := loadClient(context.Background())
	if err != nil {
		return myStruct{}, err
	}

	return myStruct{}, nil
}

type DBHandler interface {
	GetDevice(ctx context.Context, deviceID string) (model.Device, error)
	InsertDevice(ctx context.Context, device model.Device) error
}

func (m myStruct) GetDevice(ctx context.Context, deviceID string) (model.Device, error) {
	/* hamin implementation*/
	return model.Device{}, nil
}
func (m myStruct) InsertDevice(ctx context.Context, device model.Device) error {
	/* hamin implementation*/
	return nil
}

// Get One Row
func GetDevice(ctx context.Context, deviceID string) (model.Device, error) {
	out, err := dynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(config.DynamodbDeviceDB),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: deviceID},
		},
	})
	if err != nil {
		return model.Device{}, err
	}

	var device model.Device
	err = attributevalue.UnmarshalMap(out.Item, &device)
	if err != nil {
		return model.Device{}, err
	}

	return device, nil
}

func InsertDevice(ctx context.Context, device model.Device) error {
	deviceMap, err := attributevalue.MarshalMap(device)
	if err != nil {
		return err
	}

	_, err = dynamodbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(config.DynamodbDeviceDB),
		Item:      deviceMap,
	})
	if err != nil {
		return err
	}

	return nil
}
