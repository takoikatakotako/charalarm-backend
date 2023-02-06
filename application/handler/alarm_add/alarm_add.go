package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/response"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/service"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// authorizationHeader := event.Headers["Authorization"]

	// fmt.Println("-------")
	// fmt.Println(ctx)
	// fmt.Println(event)
	// fmt.Println(authorizationHeader)
	// fmt.Println("-------")

	// userID, authToken, err := auth.Basic(authorizationHeader)
	// if err != nil {
	// 	return failureResponse()
	// }

	// request := request.AddPushTokenRequest{}
	// body := event.Body
	// err = json.Unmarshal([]byte(body), &request)
	// if err != nil {
	// 	return failureResponse()
	// }
	// pushToken := request.PushToken

	// // add push token
	// s := service.PushTokenService{
	// 	DynamoDBRepository: repository.DynamoDBRepository{},
	// 	SNSRepository:      repository.SNSRepository{},
	// }
	// err = s.AddIOSVoipPushToken(userID, authToken, pushToken)
	// if err != nil {
	// 	return failureResponse()
	// }

	// jsonBytes, _ := json.Marshal("登録完了")
	// return events.APIGatewayProxyResponse{
	// 	Body:       string(jsonBytes),
	// 	StatusCode: http.StatusOK,
	// }, nil








	
	body := event.Body
	request := entity.AnonymousAddAlarmRequest{}

	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "デコードに失敗しました",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	userID := request.UserID
	userToken := request.UserToken
	alarm := request.Alarm

	s := service.AlarmService{Repository: repository.DynamoDBRepository{}}

	if err := s.AddAlarm(userID, userToken, alarm); err != nil {
		fmt.Println(err)
		response := response.MessageResponse{Message: "アラームの追加に失敗しました。"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: http.StatusInternalServerError,
		}, nil
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
