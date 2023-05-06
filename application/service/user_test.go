package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/takoikatakotako/charalarm-backend/repository"
)

func TestInfoUser(t *testing.T) {
	// DynamoDBRepository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	userService := UserService{Repository: dynamoDBRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	ipAddress := "127.0.0.1"
	platform := "iOS"
	err := userService.Signup(userID, authToken, platform, ipAddress)
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
	// DynamoDBRepository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}

	// Service
	s := UserService{Repository: dynamoDBRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	platform := "iOS"
	ipAddress := "0.0.0.0"
	err := s.Signup(userID, authToken, platform, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := dynamoDBRepository.GetUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, authToken, getUser.AuthToken)
}

func TestWithdraw(t *testing.T) {
	// DynamoDBRepository
	dynamoDBRepository := repository.DynamoDBRepository{IsLocal: true}
	s := UserService{Repository: dynamoDBRepository}

	// ユーザー作成
	userID := uuid.New().String()
	authToken := uuid.New().String()
	platform := "iOS"
	ipAddress := "0.0.0.0"
	err := s.Signup(userID, authToken, platform, ipAddress)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	firstIsExist, err := dynamoDBRepository.IsExistUser(userID)
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
	secondIsExist, err := dynamoDBRepository.IsExistUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, secondIsExist, false)
}
