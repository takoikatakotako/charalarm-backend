package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/handler"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/service"
	"net/http"
)

func Handler(ctx context.Context, event events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	s := service.WorkerService{
		SNSRepository: repository.SNSRepository{},
		SQSRepository: repository.SQSRepository{},
	}

	for _, sqsMessage := range event.Records {
		// Decode
		req := entity.IOSVoIPPushAlarmInfoSQSMessage{}
		err := json.Unmarshal([]byte(sqsMessage.Body), &req)
		if err != nil {
			// Decode失敗のためデッドレターキューに送信
			err = s.SendMessageToDeadLetter(sqsMessage.Body)
			if err == nil {
				continue
			}
			// デッドレターキューに送信にも失敗した場合
			return handler.FailureResponse(http.StatusInternalServerError, "Fail")
		}

		// メッセージを取得して処理する
		err = s.PublishPlatformApplication(req)
		if err == nil {
			continue
		}

		// デッドレターキューに送信にも失敗した場合
		return handler.FailureResponse(http.StatusInternalServerError, "Fail")
	}

	return events.APIGatewayProxyResponse{
		Body:       "Success",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
