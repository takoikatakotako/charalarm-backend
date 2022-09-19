package main

import (
	"charalarm/model"
	"charalarm/repository"
	"charalarm/entity"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := event.Body
	request := entity.AnonymousAddAlarmRequest{}

	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string("デコードに失敗しました"),
			StatusCode: 500,
		}, nil
	}

	userID := request.UserID
	userToken := request.UserToken
	alarm   := request.Alarm

	model := model.AlarmAdd{Repository: repository.DynamoDBRepository{}}
	err := model.AddAlarm(userID, userToken, alarm)
	if err != nil {
		fmt.Println(err)
		response := entity.MessageResponse{Message: "アラームの追加に失敗しました。"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: 500,
		}, nil
	}

	response := entity.MessageResponse{Message: "アラーム追加完了!"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
