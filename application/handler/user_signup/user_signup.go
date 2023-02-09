package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/takoikatakotako/charalarm-backend/handler"
	"github.com/takoikatakotako/charalarm-backend/message"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/request"
	"github.com/takoikatakotako/charalarm-backend/response"
	"github.com/takoikatakotako/charalarm-backend/service"
)

func Handler(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := event.Body
	request := request.UserSignUp{}

	// Decode Body
	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return handler.FailureResponse(http.StatusBadRequest, message.INVALID_REQUEST_PARAMETER)
	}

	// Get Parameters
	userID := request.UserID
	userToken := request.UserToken
	ipAddress := event.RequestContext.Identity.SourceIP

	// Signup
	s := service.UserService{Repository: repository.DynamoDBRepository{}}
	if err := s.Signup(userID, userToken, ipAddress); err != nil {
		return handler.FailureResponse(http.StatusBadRequest, message.USER_SIGNUP_FAILURE)
	}

	// Success
	res := response.MessageResponse{Message: message.USER_SIGNUP_SUCCESS}
	jsonBytes, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
