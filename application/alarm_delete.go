package main

import (
	"charalarm/entity"
	"charalarm/model"
	"charalarm/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := name.Body
	request := entity.AnonymousDeleteAlarmRequest{}

	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string("デコードに失敗しました"),
			StatusCode: 500,
		}, nil
	}

	userID := request.UserID
	userToken := request.UserToken
	alarmID := request.AlarmID

	model := model.AlarmDelete{Repository: repository.DynamoDBRepository{}}
	err := model.DeleteAlarm(userID, userToken, alarmID)
	if err != nil {
		fmt.Println(err)
		response := entity.MessageResponse{Message: "アラームの削除に失敗しました。"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: 500,
		}, nil
	}

	response := entity.MessageResponse{Message: "アラーム削除完了!"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
