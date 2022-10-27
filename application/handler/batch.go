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

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// 現在時刻取得
	t := time.Now()
	hour := t.Hour()
	minute := t.Minute()
	weekday := t.Weekday()

	fmt.Printf("hour: %d minute: %d\n", hour, minute)

	s := service.BatchService{
		DynamoDBRepository: repository.DynamoDBRepository{},
		SQSRepository:      repository.SQSRepository{},
	}
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
