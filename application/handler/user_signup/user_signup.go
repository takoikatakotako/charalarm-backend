package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
	"github.com/takoikatakotako/charalarm-backend/service"
)

func Handler(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := name.Body
	request := request.UserSignUp{}

	// Decode Body
	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string(message.FAILED_TO_DECODE),
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
		response := entity.MessageResponse{Message: "Sign Up Failure..."}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	response := entity.MessageResponse{Message: "Sign Up Success!"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
