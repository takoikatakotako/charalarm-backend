package model

import (
	"charalarm/repository"
	"testing"
	// "charalarm/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	repository := repository.DynamoDBRepository{IsLocal: true}
	model := SignupAnonymousUser{Repository: repository}

	userID := uuid.New().String()
	userToken := uuid.New().String()

	// ユーザー作成
	err := model.Signup(userID, userToken)
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
	// assert.Equal(t, userToken, getUser.UserToken)
}
