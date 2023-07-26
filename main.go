package main

import (
	"github.com/abdorreza/go-aws-challenge/device"
	"github.com/aws/aws-lambda-go/lambda"
)

/*func HandleRequest(ctx context.Context, event interface{}) (string, error) {
	fmt.Printf("%v\n", event)
	return "Hello from Lambda!", nil
}

func main() {
	lambda.Start(HandleRequest)
}*/

func main() {
	lambda.Start(device.Handler)
}

// Implemnting of AWS Lambda in GoLang (GET and POS Request)
