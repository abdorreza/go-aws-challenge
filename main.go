package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	region := "eu-central-1"
	profile := "myprofile"

	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{Region: aws.String(region),
			CredentialsChainVerboseErrors: aws.Bool(true)},
		Profile: profile,
	})
	if err != nil {
		fmt.Println(err)
	}
	svc := s3.New(sess)
	_, err = svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String("myxplbukcet"),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

/*func main() {
	lambda.Start(device.Add)
}*/
