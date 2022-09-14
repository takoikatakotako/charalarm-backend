package main

import (
	"charalarm/model"
	"charalarm/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	UserId    string `json: "userId"`
	UserToken string `json: "userToken"`
}

func Handler(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := name.Body
	request := Request{}

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

	userId := request.UserId
	userToken := request.UserToken

	model := model.SignupAnonymousUser{Repository: repository.DynamoDBRepository{}}
	model.Signup(userId, userToken)

	// jsonBytes, _ := json.Marshal(res)
	return events.APIGatewayProxyResponse{
		Body:       string("登録完了しました"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
