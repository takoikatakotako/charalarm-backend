package main

import (
	"context"
	"charalarm/entity"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := entity.MessageResponse{Message: "healthy!"}
	jsonBytes, _ := json.Marshal(response)
	
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
