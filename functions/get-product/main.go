package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func apiGatewayHandler(ctx context.Context, request *events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	count := 10

	if count == 10 {
		if count == 10 {
			if count == 10 {
				if count == 10 {
					if count == 10 {
						if count == 10 {
							if count == 10 {
								if count == 10 {
									if count == 10 {
										if count == 10 {
											if count == 10 {
												if count == 10 {
													if count == 10 {
														if count == 10 {
															if count == 10 {
																if count == 10 {
																	fmt.Println("Hello world!")
																}
															}

														}
													}
												}
											}

										}
									}
								}
							}

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
