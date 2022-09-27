package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity"
	charalarm_error "github.com/takoikatakotako/charalarm-backend/error"
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

	// Decode Body
	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string(charalarm_error.FAILED_TO_DECODE_REQUEST_BODY),
			StatusCode: 500,
		}, nil
	}

	// Get Parameters
	userID := request.UserID
	userToken := request.UserToken

	// Signup
	s := service.AnonymousUserService{Repository: repository.DynamoDBRepository{}}

	if err := s.Signup(userID, userToken); err != nil {
		fmt.Println(err)
		response := entity.MessageResponse{Message: "登録失敗しました"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	response := entity.MessageResponse{Message: "登録完了しました"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
