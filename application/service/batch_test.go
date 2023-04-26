package service

import (
	"github.com/google/uuid"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"testing"
)

func TestBatchService_QueryDynamoDBAndSendMessage(t *testing.T) {
	// Repository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}

	// Service
	userService := UserService{Repository: dynamoDBRepository}
	
	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	const ipAddress = "127.0.0.1"
	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
