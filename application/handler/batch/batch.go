package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity/response"
	"github.com/takoikatakotako/charalarm-backend/repository/dynamodb"
	"github.com/takoikatakotako/charalarm-backend/repository/sqs"
	"github.com/takoikatakotako/charalarm-backend/service"
	"net/http"
	"time"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 現在時刻取得
	t := time.Now().UTC()
	hour := t.Hour()
	minute := t.Minute()
	weekday := t.Weekday()

	fmt.Printf("hour: %d minute: %d\n", hour, minute)

	dynamodbRepository := &dynamodb.DynamoDBRepository{}
	sqsRepository := &sqs.SQSRepository{}
	s := service.BatchService{
		DynamoDBRepository: dynamodbRepository,
		SQSRepository:      sqsRepository,
	}
	err := s.QueryDynamoDBAndSendMessage(hour, minute, weekday)
	if err != nil {
		fmt.Println("----------------")
		fmt.Printf("Hander: %v\n", err)
		fmt.Println("----------------")

		res := response.MessageResponse{Message: "ユーザー情報の取得に失敗しました"}
		jsonBytes, _ := json.Marshal(res)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	res := response.MessageResponse{Message: "healthy!"}
	jsonBytes, _ := json.Marshal(res)

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
