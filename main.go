package main

import (
	"github.com/abdorreza/go-aws-challenge/device"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(device.Add)
}
