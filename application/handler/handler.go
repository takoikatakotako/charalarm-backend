package handler

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/takoikatakotako/charalarm-backend/response"
)

func FailureResponse(statusCode int, message string) (events.APIGatewayProxyResponse, error) {
	response := response.MessageResponse{Message: message}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: statusCode,
	}, nil
}
