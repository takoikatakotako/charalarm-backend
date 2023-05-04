package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/database"
	"github.com/takoikatakotako/charalarm-backend/repository"
)

func TestInfoUser(t *testing.T) {
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	userService := UserService{Repository: dynamoDBRepository}

	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"

	// ユーザー作成
	err := userService.Signup(userID, authToken, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := userService.GetUser(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
}

func TestSignup(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	s := UserService{Repository: repository}

	userID := uuid.New().String()
	authToken := uuid.New().String()

	// ユーザー作成
	err := s.Signup(userID, authToken, "0.0.0.0")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := repository.GetUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, authToken, getUser.AuthToken)
}

func TestWithdraw(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	s := UserService{Repository: repository}

	userID := uuid.New().String()
	authToken := uuid.New().String()

	// ユーザー作成
	insertUser := database.User{UserID: userID, AuthToken: authToken}
	err := repository.InsertUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	firstIsExist, err := repository.IsExistUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, firstIsExist, true)

	// Withdraw
	err = s.Withdraw(userID, authToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	secondIsExist, err := repository.IsExistUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, secondIsExist, false)
}
