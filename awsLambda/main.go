package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"log"
)

func invokeLambda(imageUrl string, width int) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := lambda.NewFromConfig(cfg)

	input := &lambda.InvokeInput{
		FunctionName: aws.String("resizeImageFunction"),
		Payload:      []byte(fmt.Sprintf(`{"imageUrl": "%s", "width": %d}`, imageUrl, width)),
	}

	_, err = client.Invoke(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to invoke lambda function, %v", err)
	}
}
