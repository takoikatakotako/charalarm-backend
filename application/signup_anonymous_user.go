package main

import (
	"charalarm/model"
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

type Response struct {
	YouName string `json: "name"`
	YouLike string `json: "like"`
}

func Handler(ctx context.Context, name events.APIGatewayProxyRequest) (Response, error) {
	body := name.Body
	request := Request{}

	fmt.Println("-------")
	fmt.Println(ctx)
	fmt.Println(name)
	fmt.Println(body)
	fmt.Println("-------")

	if err := json.Unmarshal([]byte(body), &request); err != nil {
		return Response{
			YouName: "デコードにしっぱいしやした",
			YouLike: "デコードにしっぱいしやした",
		}, nil
	}

	userId := request.UserId
	userToken := request.UserToken

	model := model.SignupAnonymousUser{}
	model.Setup()
	model.Signup(userId, userToken)

	return Response{
		YouName: fmt.Sprintf("UserID %s です。", request.UserId),
		YouLike: fmt.Sprintf("UserToken %s です", request.UserToken),
	}, nil
}

func main() {
	lambda.Start(Handler)
}
