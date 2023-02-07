package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/handler"
	"github.com/takoikatakotako/charalarm-backend/auth"
	"github.com/takoikatakotako/charalarm-backend/request"
	"github.com/takoikatakotako/charalarm-backend/response"
	"github.com/takoikatakotako/charalarm-backend/repository"
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

	body := event.Body
	request := request.AddAlarmRequest{}
	err = json.Unmarshal([]byte(body), &request)
	if err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, "xxxx")
	}
	alarm := request.Alarm

	s := service.AlarmService{Repository: repository.DynamoDBRepository{}}
	if err := s.AddAlarm(userID, authToken, alarm); err != nil {
		return handler.FailureResponse(http.StatusInternalServerError, "xxxx")
	}

	response := response.MessageResponse{Message: "アラーム追加完了!"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
