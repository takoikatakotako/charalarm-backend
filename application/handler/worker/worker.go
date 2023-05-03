package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/service"
	"net/http"
)

func Handler(ctx context.Context, event events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	s := service.WorkerService{
		SNSRepository: repository.SNSRepository{},
		SQSRepository: repository.SQSRepository{},
	}

	fmt.Println("--------")
	for _, message := range event.Records {
		// メッセージを取得して処理する
		err := s.PublishPlatformApplication(message.Body)
		if err == nil {
			continue
		}

		fmt.Println(err)

		// エラーの場合はデッドレターキューに格納する
		err = s.SendMessageToDeadLetter(message.Body)
		fmt.Println(err)
		return events.APIGatewayProxyResponse{
			Body:       string("Fail"),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	fmt.Println("--------")

	return events.APIGatewayProxyResponse{
		Body:       string("Success"),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
