package main

import (
	"charalarm/model"
	"fmt"
	// "github.com/aws/aws-lambda-go/lambda"
)


func main() {
	// lambda.Start(hello)
	model := model.GetAnonymousUser{}
	model.Setup()
	model.GetAnonymousUser("UUUUserDDDD")
	fmt.Printf("Hello!!")
}
