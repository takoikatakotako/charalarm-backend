package main

import (
	"charalarm/model"
	"fmt"
	// "github.com/aws/aws-lambda-go/lambda"
)


func main() {
	// lambda.Start(hello)
	model := model.AnonymousUserSignup{}
	model.Setup()
	model.Signup("UUUUserIDDDD", "UUUUUSERTTTOKEN")
	fmt.Printf("Hello!!")
}
