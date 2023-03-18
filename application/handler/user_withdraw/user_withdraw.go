package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/auth"
	"github.com/takoikatakotako/charalarm-backend/handler"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/response"
	"github.com/takoikatakotako/charalarm-backend/service"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authorizationHeader := event.Headers["Authorization"]

	fmt.Println("-------")
	fmt.Println(ctx)
	fmt.Println(event)
	fmt.Println(authorizationHeader)
	fmt.Println("-------")

	userID, authToken, err := auth.Basic(authorizationHeader)
	if err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, "xxxx")
	}

	// Withdraw
	s := service.UserService{Repository: repository.DynamoDBRepository{}}
	err = s.Withdraw(userID, authToken)
	if err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, "xxxx")
	}

	// Success
	res := response.MessageResponse{Message: "Withdraw Success!"}
	jsonBytes, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
