package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// 現在時刻取得
	t := time.Now()
    hour := t.Hour()
    minute := t.Minute()
	weekday := t.Weekday()

	s := service.BatchService{Repository: repository.DynamoDBRepository{}}
	anonymousUser, err := s.GetAnonymousUser(userID, userToken)
	if err != nil {
		fmt.Println(err)
		response := entity.MessageResponse{Message: "ユーザー情報の取得に失敗しました"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	jsonBytes, _ := json.Marshal(anonymousUser)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
	
	
	
	
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


