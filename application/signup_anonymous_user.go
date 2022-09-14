package main

import (
	"charalarm/error"
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

	// Decode Body
	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       string(charalarm_error.FAILED_TO_DECODE_REQUEST_BODY),
			StatusCode: 500,
		}, nil
	}

	// Get Parameters
	userId := request.UserId
	userToken := request.UserToken

	// Signup
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
