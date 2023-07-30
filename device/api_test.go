package device

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/abdorreza/go-aws-challenge/model"
	"github.com/aws/aws-lambda-go/events"
)

type mockDbClient struct {
	err    error
	device model.Device
	id     string
	t      *testing.T
}

func (m mockDbClient) GetDevice(ctx context.Context, deviceID string) (model.Device, error) {
	if deviceID != m.id {
		m.t.Errorf("Expected different id, got %v", deviceID)
	}

	return m.device, m.err
}

func (m mockDbClient) InsertDevice(ctx context.Context, device model.Device) error {
	if device != m.device {
		m.t.Errorf("Expected different device, got %v", device)
	}

	return m.err
}

func TestGet(t *testing.T) {
	deviceId := "12"
	device := model.Device{
		Id:          "12",
		DeviceModel: "432",
		Name:        "trx",
		Note:        "not13",
		Serial:      "t56",
	}
	mockDbHanler := mockDbClient{
		err:    nil,
		id:     deviceId,
		device: device,
		t:      t,
	}
	dbHandler = mockDbHanler

	request := events.APIGatewayProxyRequest{}
	request.PathParameters = make(map[string]string)
	request.PathParameters["id"] = deviceId

	response, err := Get(context.Background(), request)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	var unmarshalledDevice model.Device
	err = json.Unmarshal([]byte(response.Body), &unmarshalledDevice)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if unmarshalledDevice != device {
		t.Errorf("Expected different device, got %v", unmarshalledDevice)
	}
	if response.StatusCode != 200 {
		t.Errorf("Expected 200 status code, got %v", response.StatusCode)
	}
}

func TestGetDBFailure(t *testing.T) {
	deviceId := "12"
	device := model.Device{
		Id:          "12",
		DeviceModel: "432",
		Name:        "trx",
		Note:        "not13",
		Serial:      "t56",
	}
	dbErr := errors.New("DB error")
	mockDbHanler := mockDbClient{
		err:    dbErr,
		id:     deviceId,
		device: device,
		t:      t,
	}
	dbHandler = mockDbHanler

	request := events.APIGatewayProxyRequest{}
	request.PathParameters = make(map[string]string)
	request.PathParameters["id"] = deviceId

	response, err := Get(context.Background(), request)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if response.StatusCode != 500 {
		t.Errorf("Expected 500 status code, got %v", response.StatusCode)
	}
}

func TestGetDBNotFound(t *testing.T) {
	deviceId := "12"
	var device model.Device
	mockDbHanler := mockDbClient{
		err:    nil,
		id:     deviceId,
		device: device,
		t:      t,
	}
	dbHandler = mockDbHanler

	request := events.APIGatewayProxyRequest{}
	request.PathParameters = make(map[string]string)
	request.PathParameters["id"] = deviceId

	response, err := Get(context.Background(), request)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if response.StatusCode != 404 {
		t.Errorf("Expected 404 status code, got %v", response.StatusCode)
	}
}

func TestAdd(t *testing.T) {
	deviceId := "12"
	device := model.Device{
		Id:          "12",
		DeviceModel: "432",
		Name:        "trx",
		Note:        "not13",
		Serial:      "t56",
	}
	mockDbHanler := mockDbClient{
		err:    nil,
		id:     deviceId,
		device: device,
		t:      t,
	}
	dbHandler = mockDbHanler

	request := events.APIGatewayProxyRequest{}

	marshalledDevice, err := json.Marshal(device)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	request.Body = string(marshalledDevice)

	response, err := Add(context.Background(), request)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if response.StatusCode != 201 {
		t.Errorf("Expected 201 status code, got %v", response.StatusCode)
	}
}

func TestAddDBFailure(t *testing.T) {
	deviceId := "12"
	dbErr := errors.New("DB Error")
	device := model.Device{
		Id:          "12",
		DeviceModel: "432",
		Name:        "trx",
		Note:        "not13",
		Serial:      "t56",
	}
	mockDbHanler := mockDbClient{
		err:    dbErr,
		id:     deviceId,
		device: device,
		t:      t,
	}
	dbHandler = mockDbHanler

	request := events.APIGatewayProxyRequest{}

	marshalledDevice, err := json.Marshal(device)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	request.Body = string(marshalledDevice)

	response, err := Add(context.Background(), request)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}
	if response.StatusCode != 500 {
		t.Errorf("Expected 500 status code, got %v", response.StatusCode)
	}
}
