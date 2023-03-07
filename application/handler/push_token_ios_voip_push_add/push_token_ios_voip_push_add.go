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
		return handler.FailureResponse(http.StatusInternalServerError, "xxxx")
	}

	request := request.AddPushTokenRequest{}
	body := event.Body
	err = json.Unmarshal([]byte(body), &request)
	if err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, "xxxx")
	}
	pushToken := request.PushToken

	// add push token
	s := service.PushTokenService{
		DynamoDBRepository: repository.DynamoDBRepository{},
		SNSRepository:      repository.SNSRepository{},
	}
	err = s.AddIOSVoipPushToken(userID, authToken, pushToken)
	if err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, "xxxx")
	}

	jsonBytes, _ := json.Marshal("登録完了")
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
