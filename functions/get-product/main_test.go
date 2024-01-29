package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/require"
)

func TestHandler(t *testing.T) {
	c := require.New(t)

	request := &events.APIGatewayProxyRequest{}

	response, err := apiGatewayHandler(context.Background(), request)
	c.NoError(err)
	c.NotNil(response)
}
