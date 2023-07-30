package device

import (
	"context"
	"encoding/json"

	"github.com/abdorreza/go-aws-challenge/db"
	"github.com/abdorreza/go-aws-challenge/model"
	"github.com/aws/aws-lambda-go/events"
)

var dbHandler db.DBHandler

func init() {
	var err error
	dbHandler, err = db.NewDynamodb()
	if err != nil {
		panic(err)
	}
}

func Get(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Retrieve the "deviceId" path parameter from the request
	deviceId := request.PathParameters["id"]

	device, err := dbHandler.GetDevice(ctx, deviceId)

	// Any Problem
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			//Body:       "HTTP 500 Internal Server Error",
			Body: err.Error(),
		}, nil
	}

	// Not Found
	if device.Id == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: 404,
			Body:       "HTTP 404 Not Found",
		}, nil
	}

	deviceJson, err := json.Marshal(device)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "HTTP 500 Internal Server Error",
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(deviceJson),
	}, nil
}

func Add(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var device model.Device
	err := json.Unmarshal([]byte(request.Body), &device)
	// Bad request
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "HTTP 400 Bad Request",
		}, nil
	}

	err = dbHandler.InsertDevice(ctx, device)
	// Server side error
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "HTTP 500 Internal Server Error",
		}, nil
	}

	// Success
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
		Body:       "HTTP 201 Created",
	}, nil
}
