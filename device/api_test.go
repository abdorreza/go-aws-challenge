package device

import (
	"context"
	"fmt"
	"testing"

	"github.com/abdorreza/go-aws-challenge/db"
	"github.com/aws/aws-lambda-go/events"
)

func TestGet(t *testing.T) {
	// Mocking the context and request for testing
	ctx := context.Background()
	request := events.APIGatewayProxyRequest{
		PathParameters: map[string]string{
			"id": "some_device_id",
		},
	}

	// Test case 1: Device found
	// Mocking the db.GetDevice function to return a device
	mockDevice := &Device{Id: "some_device_id", Name: "Test Device"}
	db.GetDevice = func(ctx context.Context, deviceId string) (*Device, error) {
		return mockDevice, nil
	}

	response, err := Get(ctx, request)

	// Verify the response and error
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, but got %d", response.StatusCode)
	}
	expectedBody := `{"Id":"some_device_id","Name":"Test Device"}`
	if response.Body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, response.Body)
	}

	// Test case 2: Device not found
	// Mocking the db.GetDevice function to return an error
	db.GetDevice = func(ctx context.Context, deviceId string) (*Device, error) {
		return nil, fmt.Errorf("Device not found")
	}

	response, err = Get(ctx, request)

	// Verify the response and error
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if response.StatusCode != 404 {
		t.Errorf("Expected status code 404, but got %d", response.StatusCode)
	}
	expectedBody = "Device not found"
	if response.Body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, response.Body)
	}

	// Test case 3: Error from db.GetDevice
	// Mocking the db.GetDevice function to return an error
	db.GetDevice = func(ctx context.Context, deviceId string) (*Device, error) {
		return nil, fmt.Errorf("Some error")
	}

	response, err = Get(ctx, request)

	// Verify the response and error
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if response.StatusCode != 500 {
		t.Errorf("Expected status code 500, but got %d", response.StatusCode)
	}
	expectedBody = "Some error"
	if response.Body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, response.Body)
	}
}


func TestAdd(t *testing.T) {
	// Mocking the context and request for testing
	ctx := context.Background()
	requestWithValidData := events.APIGatewayProxyRequest{
		Body: `{"Id":"some_device_id","Name":"Test Device"}`,
	}
	requestWithInvalidData := events.APIGatewayProxyRequest{
		Body: "invalid JSON data",
	}

	// Test case 1: Valid data, should return 201 Created
	// Mocking the db.InsertDevice function to succeed
	db.InsertDevice = func(ctx context.Context, device model.Device) error {
		return nil
	}

	response, err := Add(ctx, requestWithValidData)

	// Verify the response and error
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if response.StatusCode != 201 {
		t.Errorf("Expected status code 201, but got %d", response.StatusCode)
	}
	expectedBody := "HTTP 201 Created"
	if response.Body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, response.Body)
	}

	// Test case 2: Invalid JSON data, should return 400 Bad Request
	response, err = Add(ctx, requestWithInvalidData)

	// Verify the response and error
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if response.StatusCode != 400 {
		t.Errorf("Expected status code 400, but got %d", response.StatusCode)
	}
	expectedBody = "invalid character 'i' looking for beginning of value"
	if response.Body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, response.Body)
	}

	// Test case 3: Error from db.InsertDevice, should return 500 Internal Server Error
	// Mocking the db.InsertDevice function to return an error
	db.InsertDevice = func(ctx context.Context, device model.Device) error {
		return fmt.Errorf("Some error")
	}

	response, err = Add(ctx, requestWithValidData)

	// Verify the response and error
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
	if response.StatusCode != 500 {
		t.Errorf("Expected status code 500, but got %d", response.StatusCode)
	}
	expectedBody = "Some error"
	if response.Body != expectedBody {
		t.Errorf("Expected body '%s', but got '%s'", expectedBody, response.Body)
	}
}
