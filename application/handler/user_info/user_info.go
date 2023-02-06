package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/auth"
	"github.com/takoikatakotako/charalarm-backend/response"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/service"
	"net/http"
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
		return failureResponse()
	}

	// info
	s := service.UserService{Repository: repository.DynamoDBRepository{}}
	anonymousUser, err := s.GetUser(userID, authToken)
	if err != nil {
		return failureResponse()
	}

	// Success
	jsonBytes, _ := json.Marshal(anonymousUser)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func failureResponse() (events.APIGatewayProxyResponse, error) {
	response := response.MessageResponse{Message: "ユーザー情報の取得に失敗しました"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusInternalServerError,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
