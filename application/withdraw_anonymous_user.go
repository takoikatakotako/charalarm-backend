package main

import (
	"charalarm/model"
	"charalarm/repository"
	"charalarm/entity"
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

	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string("デコードに失敗しました"),
			StatusCode: 500,
		}, nil
	}

	userID := request.UserID
	userToken := request.UserToken

	// Withdraw
	model := model.WithdrawAnonymousUser{Repository: repository.DynamoDBRepository{}}
	err := model.Withdraw(userID, userToken)
	if err != nil {
		fmt.Println(err)
		response := entity.MessageResponse{Message: "退会失敗しました"}
		jsonBytes, _ := json.Marshal(response)
		return events.APIGatewayProxyResponse{
			Body:       string(jsonBytes),
			StatusCode: 500,
		}, nil
	}

	// jsonBytes, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		Body:       string("退会完了しました"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
