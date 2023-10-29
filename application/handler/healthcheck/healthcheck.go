package main

import (
	"context"
	"encoding/json"
	"github.com/takoikatakotako/charalarm-backend/entity/response"
	"github.com/takoikatakotako/charalarm-backend/util/message"
	"log/slog"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, name events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res := response.MessageResponse{Message: message.Healthy}
	jsonBytes, _ := json.Marshal(res)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello structured log", "name", "blue", "No", 5)

	logger.Error("This is Error!!")
	logger.Warn("This is Warn!")
	logger.Info("This is Info!")

	return events.APIGatewayProxyResponse{
		Body:       string(jsonBytes),
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
