package db

import (
	"context"

	"github.com/abdorreza/go-aws-challenge/model"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type dynamodbHandler interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

type DBHandler interface {
	GetDevice(ctx context.Context, deviceID string) (model.Device, error)
	InsertDevice(ctx context.Context, device model.Device) error
}
