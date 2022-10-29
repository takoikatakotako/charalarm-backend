package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/service"
	"net/http"
	"time"
	"fmt"
)

func Handler(ctx context.Context, event events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	s := service.WorkerhService{
		SNSRepository: repository.SNSRepository{},
		SQSRepository: repository.SQSRepository{},
	}


	fmt.Println("--------")

    for _, message := range event.Records {
		// メッセージを取得して処理する
		err := s.SendAlarmInfoMessage

		// エラーの場合はデッドレターキューに格納する


	}
	fmt.Println("--------")




    for record in event['Records']:
        body['messages'].append(record["body"])




	err := s.QueryDynamoDBAndSendMessage(hour, minute, weekday)
	if err != nil {
		response := entity.MessageResponse{Message: "ユーザー情報の取得に失敗しました"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	response := entity.MessageResponse{Message: "healthy!"}
	jsonBytes, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
