package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func apiGatewayHandler(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	count := 10

	for a := 0; a < count; a++ {
		for b := 0; b < count; b++ {
			for c := 0; c < count; c++ {
				for d := 0; d < count; d++ {
					for e := 0; e < count; e++ {
						for f := 0; f < count; f++ {
							fmt.Println("Hello world!")
						}
					}
				}
			}
		}
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Hello world",
	}, nil
}

func main() {
	lambda.Start(apiGatewayHandler)
}
