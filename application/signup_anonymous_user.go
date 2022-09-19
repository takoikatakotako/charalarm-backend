package main

import (
	"charalarm/entity"
	"charalarm/error"
	"charalarm/model"
	"charalarm/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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
	model := model.SignupAnonymousUser{Repository: repository.DynamoDBRepository{}}
	err := model.Signup(userID, userToken)
	if err != nil {
		fmt.Println(err)
		response := entity.MessageResponse{Message: "登録失敗しました"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: 500,
		}, nil
	}

	response := entity.MessageResponse{Message: "登録完了しました"}
	jsonBytes, _ := json.Marshal(response)
	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
