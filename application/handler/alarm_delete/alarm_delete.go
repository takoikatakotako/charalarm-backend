package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity"
	repository "github.com/takoikatakotako/charalarm-backend/repository/aws"
	"github.com/takoikatakotako/charalarm-backend/service"
)

func Handler(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := name.Body
	request := entity.AnonymousDeleteAlarmRequest{}

	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "デコードに失敗しました",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	userID := request.UserID
	userToken := request.UserToken
	alarmID := request.AlarmID

	s := service.AlarmService{Repository: repository.DynamoDBRepository{}}
	if err := s.DeleteAlarm(userID, userToken, alarmID); err != nil {
		fmt.Println(err)
		response := entity.MessageResponse{Message: "アラームの削除に失敗しました。"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	response := entity.MessageResponse{Message: "アラーム削除完了!"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
