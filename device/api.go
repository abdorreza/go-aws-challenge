package device

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

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
