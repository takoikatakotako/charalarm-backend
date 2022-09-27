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
	request := entity.AnonymousUserRequest{}

	fmt.Println("-------")
	fmt.Println(ctx)
	fmt.Println(name)
	fmt.Println(body)
	fmt.Println("-------")

	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "デコードに失敗しました",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	userID := request.UserID
	userToken := request.UserToken

	s := service.AnonymousUserService{Repository: repository.DynamoDBRepository{}}
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
}

func main() {
	lambda.Start(Handler)
}
