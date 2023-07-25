package main

import (
	"github.com/abdorreza/go-aws-challenge/device"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(device.Handler)
}

// Implemnting of AWS Lambda in GoLang (GET and POS Request)
