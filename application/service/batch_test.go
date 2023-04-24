package service

import (
	"github.com/google/uuid"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"testing"
)

func TestBatch(t *testing.T) {
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
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
