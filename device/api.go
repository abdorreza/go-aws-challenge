package device

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

func Get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Retrieve the "deviceId" path parameter from the request
	deviceId := request.PathParameters["id"]

	// Your logic to fetch device data based on "deviceId"
	// Replace the below sample response with your actual data retrieval code
	deviceData := fmt.Sprintf(`{"deviceId": "%s", "name": "Sensor", "note": "Testing a sensor."}`, deviceId)

	// Create the HTT
	// Create the HTTP response with status code 200 and the device data
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deviceData,
	}, nil
}

func Add(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Retrieve the "deviceId" path parameter from the request
	fmt.Println("*******************************************************")
	deviceId := request.PathParameters["id"]

	// Your logic to fetch device data based on "deviceId"
	// Replace the below sample response with your actual data retrieval code
	deviceData := fmt.Sprintf(`{"deviceId": "%s", "name": "Sensor", "note": "Testing a sensor."}`, deviceId)

	// Create the HTT
	// Create the HTTP response with status code 200 and the device data
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       deviceData,
	}, nil
}
