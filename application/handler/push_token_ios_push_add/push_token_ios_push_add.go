package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/takoikatakotako/charalarm-backend/handler"
	"github.com/takoikatakotako/charalarm-backend/message"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/auth"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
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
		return handler.FailureResponse(http.StatusInternalServerError, message.AuthenticationFailure)
	}

	req := request.AddPushTokenRequest{}
	body := event.Body
	err = json.Unmarshal([]byte(body), &req)
	if err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, message.InvalidValue)
	}
	pushToken := req.PushToken

	// add push token
	s := service.PushTokenService{
		DynamoDBRepository: &repository.DynamoDBRepository{},
		SNSRepository:      repository.SNSRepository{},
	}
	err = s.AddIOSPushToken(userID, authToken, pushToken)
	if err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, message.UserUpdateFailure)
	}

	jsonBytes, _ := json.Marshal(message.UserUpdateSuccess)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
