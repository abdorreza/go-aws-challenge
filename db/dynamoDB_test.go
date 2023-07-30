package db

import (
	"context"
	"errors"
	"sync"
	"testing"

	"github.com/abdorreza/go-aws-challenge/model"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type mockDynamoDBClient struct {
	err    error
	device model.Device
	id     string
	t      *testing.T
}

func (m *mockDynamoDBClient) GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error) {
	id := params.Key["id"]
	if id.(*types.AttributeValueMemberS).Value != m.id {
		m.t.Error("Get Item id is not expected")
	}
	marshalDevice, err := attributevalue.MarshalMap(m.device)
	if err != nil {
		m.t.Error("MarshalMap failed")
	}
	return &dynamodb.GetItemOutput{Item: marshalDevice}, m.err
}

func (m *mockDynamoDBClient) PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error) {
	var unmarshalDevice model.Device
	err := attributevalue.UnmarshalMap(params.Item, &unmarshalDevice)
	if err != nil {
		m.t.Error("UnmarshalMap failed")
	}
	if unmarshalDevice != m.device {
		m.t.Error("Put Item device invalid")
	}
	return &dynamodb.PutItemOutput{}, m.err
}

func TestLoadClient(t *testing.T) {
	once = sync.Once{}
	dynamodbClient = nil

	err := loadClient(context.Background())
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if dynamodbClient == nil {
		t.Error("Expected dynamodbClient not to be nil")
	}
}

func TestGetDevice(t *testing.T) {
	device := model.Device{
		Id:          "12",
		DeviceModel: "432",
		Name:        "trx",
		Note:        "not13",
		Serial:      "t56",
	}
	dynamodbClient = &mockDynamoDBClient{
		err:    nil,
		id:     device.Id,
		device: device,
		t:      t,
	}

	myStruct1 := dynamodbStruct{}
	fetchedDevice, err := myStruct1.GetDevice(context.Background(), device.Id)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if fetchedDevice != device {
		t.Errorf("Expected different device, got %v", fetchedDevice)
	}
}

func TestGetDeviceAWSFailure(t *testing.T) {
	id := "12"
	awsErr := errors.New("AWS Error")
	dynamodbClient = &mockDynamoDBClient{
		err:    awsErr,
		id:     id,
		device: model.Device{},
		t:      t,
	}

	myStruct1 := dynamodbStruct{}
	fetchedDevice, err := myStruct1.GetDevice(context.Background(), id)
	if err != awsErr {
		t.Errorf("Expected different error, got %v", err)
	}

	emptyDevice := model.Device{}
	if fetchedDevice != emptyDevice {
		t.Errorf("Expected empty device, got %v", fetchedDevice)
	}
}

func TestInsertDevice(t *testing.T) {
	device := model.Device{
		Id:          "12",
		DeviceModel: "432",
		Name:        "trx",
		Note:        "not13",
		Serial:      "t56",
	}
	dynamodbClient = &mockDynamoDBClient{
		err:    nil,
		id:     device.Id,
		device: device,
		t:      t,
	}

	myStruct1 := dynamodbStruct{}
	err := myStruct1.InsertDevice(context.Background(), device)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
}

func TestInsertDeviceAWSFailure(t *testing.T) {
	awsErr := errors.New("AWS ERROR")
	device := model.Device{
		Id:          "12",
		DeviceModel: "432",
		Name:        "trx",
		Note:        "not13",
		Serial:      "t56",
	}
	dynamodbClient = &mockDynamoDBClient{
		err:    awsErr,
		id:     device.Id,
		device: device,
		t:      t,
	}

	myStruct1 := dynamodbStruct{}
	err := myStruct1.InsertDevice(context.Background(), device)
	if err != awsErr {
		t.Errorf("Expected different error, got %v", err)
	}
}
