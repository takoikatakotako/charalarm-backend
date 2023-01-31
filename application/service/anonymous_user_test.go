package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	// "github.com/takoikatakotako/charalarm-backend/entity"
	"github.com/takoikatakotako/charalarm-backend/repository"
	"github.com/takoikatakotako/charalarm-backend/database"
)

func TestInfoUser(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	s := AnonymousUserService{Repository: repository}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// ユーザー作成
	insertUser := database.User{UserID: userID, UserToken: userToken}
	err := repository.InsertAnonymousUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := s.GetAnonymousUser(userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, userToken, getUser.UserToken)
}

func TestSignup(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	s := AnonymousUserService{Repository: repository}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// ユーザー作成
	err := s.Signup(userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// ユーザー取得
	getUser, err := repository.GetAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, userID, getUser.UserID)
	assert.Equal(t, userToken, getUser.UserToken)
}

func TestWithdraw(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	s := AnonymousUserService{Repository: repository}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// ユーザー作成
	insertUser := database.User{UserID: userID, UserToken: userToken}
	err := repository.InsertAnonymousUser(insertUser)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	firstIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, firstIsExist, true)

	// Withdraw
	err = s.Withdraw(userID, userToken)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// IsExist
	secondIsExist, err := repository.IsExistAnonymousUser(userID)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	// Assert
	assert.Equal(t, secondIsExist, false)
}
