package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/takoikatakotako/charalarm-backend/entity/response"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/entity/request"
	"github.com/takoikatakotako/charalarm-backend/handler"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/service"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := event.Body

	fmt.Println("-------")
	fmt.Println(ctx)
	fmt.Println(event)
	fmt.Println(body)
	fmt.Println("-------")

	// Decode Body
	req := request.UserSignUp{}
	if err := json.Unmarshal([]byte(body), &req); err != nil {
		return handler.FailureResponse(http.StatusBadRequest, message.InvalidRequestParameter)
	}

	// Get Parameters
	userID := req.UserID
	authToken := req.AuthToken
	platform := req.Platform
	ipAddress := event.RequestContext.Identity.SourceIP

	// Signup
	s := service.UserService{DynamoDBRepository: &repository.DynamoDBRepository{}}
	if err := s.Signup(userID, authToken, platform, ipAddress); err != nil {
		fmt.Println(err)
		return handler.FailureResponse(http.StatusBadRequest, message.UserSignupFailure)
	}

	// Success
	res := response.MessageResponse{Message: message.UserSignupSuccess}
	jsonBytes, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
